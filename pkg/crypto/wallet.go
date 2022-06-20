package crypto

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MPCWallet struct {
	pool      *pool.Pool
	ID        party.ID
	Configs   map[party.ID]*cmp.Config
	Network   *Network
	Threshold int
}

// Generate a new ECDSA private key shared among all the given participants.
func Generate(options ...WalletOption) (*MPCWallet, error) {
	opt := defaultConfig()
	wallet := opt.Apply(options...)

	var wg sync.WaitGroup
	for _, id := range opt.participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			defer wg.Done()
			conf, err := cmpKeygen(id, opt.participants, wallet.Network, opt.threshold, pl)
			if err != nil {
				return
			}
			wallet.Configs[conf.ID] = conf
		}(id)
	}
	wg.Wait()
	return wallet, nil
}

// Returns the Bech32 representation of the given party.
func (w *MPCWallet) Address(id ...party.ID) (string, error) {
	pub, err := w.PublicKeyProto()
	if err != nil {
		return "", err
	}

	str, err := bech32.ConvertAndEncode("snr", pub.Address().Bytes())
	if err != nil {
		return "", err
	}
	return str, nil
}

// Config returns the configuration of this wallet.
func (w *MPCWallet) Config() *cmp.Config {
	return w.Configs[w.ID]
}

// DID Returns the DID Address of the Wallet. When partyId is provided, it returns the DID of the given party. Only the first party in the wallet can create a DID.
func (w *MPCWallet) DID(party ...party.ID) (*did.DID, error) {
	if len(party) == 0 {
		addr, err := w.Address()
		if err != nil {
			return nil, err
		}
		return did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(addr, "snr")))
	} else if len(party) == 1 {
		id := party[0]
		if !w.Config().PartyIDs().Contains(id) {
			return nil, fmt.Errorf("party %s is not a member of the wallet", id)
		}

		baseDid, err := w.DID()
		if err != nil {
			return nil, err
		}
		return did.ParseDID(fmt.Sprintf("%s#%s", baseDid, id))
	}
	return nil, fmt.Errorf("invalid number of arguments")
}

// CreateDIDDocument creates a DID Document for the given party.
func (w *MPCWallet) DIDDocument() (did.Document, error) {
	// Get the DID of the wallet
	baseDid, err := w.DID()
	if err != nil {
		return nil, err
	}

	// Create the DID Document
	doc, err := did.NewDocument(baseDid.String())
	if err != nil {
		return nil, err
	}

	// Get ALL the VerificationMethods of the wallet.
	vmsAll := make([]*did.VerificationMethod, 0)
	for _, id := range w.Configs[w.ID].PartyIDs() {
		vm, err := w.GetVerificationMethod(id)
		if err != nil {
			return nil, err
		}
		vmsAll = append(vmsAll, vm)
	}

	for _, vm := range vmsAll {
		doc.AddAuthenticationMethod(vm)
	}

	if len(doc.GetAuthenticationMethods()) != len(vmsAll) {
		return nil, fmt.Errorf("failed to add all verification methods to DID Document")
	}
	return doc, nil
}

// GetSigners returns the list of signers for the given message.
func (w *MPCWallet) GetSigners() party.IDSlice {
	signers := party.IDSlice([]party.ID{"dsc", "psk"})
	// signers := w.Configs[w.ID].PartyIDs()[:w.Threshold+1]
	if !signers.Contains(w.ID) {
		w.Network.Quit(w.ID)
		return nil
	}
	return party.NewIDSlice(signers)
}

// GetVerificationMethod returns the VerificationMethod for the given party.
func (w *MPCWallet) GetVerificationMethod(id party.ID) (*did.VerificationMethod, error) {
	// Get the DID of the wallet
	baseDid, err := w.DID()
	if err != nil {
		return nil, err
	}

	// Get the DID of the party.
	vmdid, err := w.DID(id)
	if err != nil {
		return nil, err
	}

	// Get base58 encoded public key.
	pub, err := w.PublicKeyBase58()
	if err != nil {
		return nil, err
	}

	// Return the shares VerificationMethod
	return &did.VerificationMethod{
		ID:              *vmdid,
		Type:            ssi.ECDSASECP256K1VerificationKey2019,
		Controller:      *baseDid,
		PublicKeyBase58: pub,
	}, nil
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKey() ([]byte, error) {
	p := w.Config().PublicPoint().(*curve.Secp256k1Point)
	buf, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	// Check length of the public key.
	if len(buf) != 33 {
		return nil, fmt.Errorf("invalid public key length")
	}
	return buf, nil
}

// Returns the ECDSA public key of the given party.
func (w *MPCWallet) PublicKeyBase58() (string, error) {
	pub, err := w.PublicKey()
	if err != nil {
		return "", err
	}
	return base58.Encode(pub), nil
}

func (w *MPCWallet) PublicKeyProto() (*secp256k1.PubKey, error) {
	pubBz, err := w.PublicKey()
	if err != nil {
		return nil, err
	}
	return &secp256k1.PubKey{
		Key: pubBz,
	}, nil
}

// Refreshes all shares of an existing ECDSA private key.
func (w *MPCWallet) Refresh(pl *pool.Pool) (*cmp.Config, error) {
	hRefresh, err := protocol.NewMultiHandler(cmp.Refresh(w.Configs[w.ID], pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(w.ID, hRefresh, w.Network)

	r, err := hRefresh.Result()
	if err != nil {
		return nil, err
	}
	handlerLoop(w.ID, hRefresh, w.Network)
	return r.(*cmp.Config), nil
}

// Generates an ECDSA signature for messageHash.
func (w *MPCWallet) Sign(m []byte) (*ecdsa.Signature, error) {
	var wg sync.WaitGroup
	signers := w.GetSigners()
	net := NewNetwork(signers)

	var (
		sig *ecdsa.Signature
		err error
	)

	for _, id := range signers {
		wg.Add(1)
		go func(id party.ID) {
			defer wg.Done()

			pl := pool.NewPool(0)
			defer pl.TearDown()
			if sig, err = cmpSign(w.Configs[id], m, signers, net, pl); err != nil {
				return
			}
		}(id)
	}
	wg.Wait()
	return sig, err
}

// Verifies an ECDSA signature for messageHash.
func (w *MPCWallet) Verify(m []byte, sig []byte) bool {
	edsig, err := SignatureFromBytes(sig)
	if err != nil {
		return false
	}
	mpcVerif := edsig.Verify(w.Config().PublicPoint(), m)
	return mpcVerif
}

func (w *MPCWallet) CreateInitialShards() (dscShard, pskShard, recShard string, unused []string, err error) {
	ss, e := w.serializedShards()
	if e != nil {
		err = e
		return
	}
	if len(ss) < 6 {
		err = errors.New("not enough shards")
		return
	}

	var ok bool
	// assign dscShard
	// TODO: encrypt with dsc (i.e. webauthn)
	dscShard, ok = ss["dsc"]
	if !ok {
		err = errors.New("could not find dsc shard")
		return
	}

	// assign pskShard
	// TODO: This should be encrypted with PSK
	pskShard, ok = ss["psk"]
	if !ok {
		err = errors.New("could not find psk shard")
		return
	}

	// assign recshard
	// TODO: this should be password encrypted
	recShard, ok = ss["recovery"]
	if !ok {
		err = errors.New("could not find recovery shard")
		return
	}

	// sign unused shards
	for i := 0; i < len(ss); i++ {
		u, ok := ss[fmt.Sprintf("bank%d", i+1)]
		if !ok {
			continue
		}
		sig, e := w.Sign([]byte(u))
		if e != nil {
			err = e
			return
		}
		sSig, e := SerializeSignature(sig)
		if e != nil {
			err = e
			return
		}
		unused = append(unused, string(sSig))
	}
	if len(unused) == 0 {
		err = errors.New("no backup shards")
	}
	return
}

func (w *MPCWallet) serializedShards() (map[string]string, error) {
	deviceShards := make(map[string]string)
	for k, c := range w.Configs {
		b, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		deviceShards[string(k)] = base64.StdEncoding.EncodeToString([]byte(b))
	}
	return deviceShards, nil
}
