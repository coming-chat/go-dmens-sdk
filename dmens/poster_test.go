package dmens

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/coming-chat/go-sui/v2/types"
	"github.com/coming-chat/wallet-SDK/core/sui"
	"github.com/stretchr/testify/require"
)

var whoami = ""
var M1 = os.Getenv("WalletSdkTestM1")
var M1Account, _ = sui.NewAccountWithMnemonic(M1)
var M1Address = M1Account.Address()

func init() {
	out, _ := exec.Command("whoami").Output()
	whoami = strings.TrimSpace(string(out))
}

func DefaultPoster(t *testing.T) *Poster {
	poster, err := NewPoster(&PosterConfig{Address: M1Address, Reviewing: true}, TestnetConfig)
	require.Nil(t, err)
	return poster
}

func TestNewPoster(t *testing.T) {
	poster := DefaultPoster(t)
	require.NotEqual(t, poster.DmensNftId, "")
}

func TestPostNote(t *testing.T) {
	poster := DefaultPoster(t)

	chain := poster.chain
	// chain := sui.NewChainWithRpcUrl("https://fullnode.testnet.sui.io")

	txn, err := poster.DmensPost("Hello world")
	require.Nil(t, err)

	fee, err := chain.EstimateTransactionFee(txn)
	require.Nil(t, err)
	t.Log(fee.Value)

	simulateCheck(t, chain, txn, true)
}

func simulateCheck(t *testing.T, chain *sui.Chain, txn *sui.Transaction, showJson bool) *types.DryRunTransactionBlockResponse {
	cli, err := chain.Client()
	require.Nil(t, err)
	resp, err := cli.DryRunTransaction(context.Background(), txn.TransactionBytes())
	require.Nil(t, err)
	require.Equal(t, resp.Effects.Data.V1.Status.Error, "")
	require.True(t, resp.Effects.Data.IsSuccess())
	if showJson {
		data, err := json.Marshal(resp)
		require.Nil(t, err)
		respStr := string(data)
		t.Log("simulate run resp: ", respStr)
	}
	return resp
}
