package client

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/third_party/types/common"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/assert"
)

func Test_FaucetCheckBalance(t *testing.T) {
	// Create Client instance and Generate wallet
	client := NewClient(mt.ClientMode_ENDPOINT_BETA)
	w, err := mpc.GenerateWallet(common.DefaultCallback())
	assert.NoError(t, err, "wallet generation succeeds")

	// Get Wallet Address
	addr, err := w.Address()
	assert.NoError(t, err, "Bech32Address successfully created")
	fmt.Println("Address:", addr)

	// Check Balance
	resp, err := client.CheckBalance(addr)
	assert.NoError(t, err, "Check Balance succeeds")
	fmt.Printf("-- Get Balances (1) --\n%+v\n", resp)

	// Request Faucet and Check Again
	err = client.RequestFaucet(addr)
	assert.NoError(t, err, "faucet request succeeds")
	resp2, err := client.CheckBalance(addr)
	assert.NoError(t, err, "Check Balance succeeds")
	fmt.Printf("-- Get Balances (2) --\n%+v\n", resp2)
}

func Test_QueryWhoIs(t *testing.T) {
	accAddr := "snr1xurfhe4cfu29k04r6rlmaqcrjzef2le46qy9rm"
	client := NewClient(mt.ClientMode_ENDPOINT_BETA)
	acc, err := client.QueryWhoIs(accAddr)
	assert.NoError(t, err, "QueryAccount succeeds")
	fmt.Printf("-- Get Account --\n%+v\n", acc)
}
