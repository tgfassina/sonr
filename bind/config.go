package sonr

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/denisbrodbeck/machineid"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	lf "github.com/sonr-io/core/internal/lifecycle"
	sl "github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	tf "github.com/sonr-io/core/internal/transfer"
)

// ^ CurrentFile returns last file in Processed Files ^ //
func (sn *Node) currentFile() *sf.ProcessedFile {
	return sn.files[len(sn.files)-1]
}

// ^ getDeviceID sets node device ID from path if Exists ^ //
func getDeviceID(connEvent *md.ConnectionRequest) error {
	// Check if ID already provided
	if connEvent.Device.Id != "" {
		return nil
	}

	// Create Device ID Path
	path := filepath.Join(connEvent.Directories.Documents, ".sonr-device-id")

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Generate ID
		id, err := machineid.ProtectedID("Sonr")
		if err != nil {
			return err
		}

		// Write ID To File
		f, err := os.Create(path)
		if err != nil {
			return err
		}

		// Defer Close
		defer f.Close()

		// Write to File
		_, err = f.WriteString(id)
		if err != nil {
			return err
		}

		// Update Device
		connEvent.Device.Id = id
		return nil
	} else {
		// @ Read Device ID Data
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Convert to String
		id := string(dat)

		// Update Device
		connEvent.Device.Id = id

		return nil
	}
}

// ^ setInfo sets node info from connEvent and host ^ //
func (sn *Node) setInfo(connEvent *md.ConnectionRequest) error {
	// Check for Host
	if sn.host == nil {
		err := errors.New("setPeer: Host has not been called")
		return err
	}

	// Set Default Properties
	sn.contact = connEvent.Contact
	sn.directories = connEvent.Directories

	// Get Device ID
	err := getDeviceID(connEvent)
	if err != nil {
		return err
	}

	// Set Peer Info
	sn.peer = &md.Peer{
		Id:         sn.host.ID().String(),
		Username:   connEvent.Username,
		Device:     connEvent.Device,
		FirstName:  connEvent.Contact.FirstName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}
	return nil
}

// ^ setConnection initializes connection protocols joins lobby and creates pubsub service ^ //
func (sn *Node) setConnection(ctx context.Context) error {
	// Create a new PubSub service using the GossipSub router
	var err error
	sn.pubSub, err = pubsub.NewGossipSub(ctx, sn.host)
	if err != nil {
		return err
	}

	// Create Callbacks
	lobCall := lf.LobbyCallbacks{CallRefresh: sn.call.OnRefreshed, CallError: sn.error, GetPeer: sn.Peer}
	transCall := lf.TransferCallbacks{CallInvited: sn.call.OnInvited, CallResponded: sn.call.OnResponded, CallReceived: sn.call.OnReceived, CallProgress: sn.call.OnProgress, CallTransmitted: sn.call.OnTransmitted, CallError: sn.error}

	// Enter Lobby
	if sn.lobby, err = sl.Join(sn.ctx, lobCall, sn.pubSub, sn.host.ID(), sn.peer, sn.olc); err != nil {
		return err
	}

	// Initialize Peer Connection
	if sn.peerConn, err = tf.Initialize(sn.host, sn.pubSub, sn.directories, sn.olc, transCall); err != nil {
		return err
	}
	return nil
}
