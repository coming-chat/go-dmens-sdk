package dmens

import (
	"os"
	"os/exec"
	"strings"
	"testing"

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
	poster, err := NewPoster(&PosterConfig{Address: M1Address}, TestnetConfig)
	require.Nil(t, err)
	return poster
}

func TestNewPoster(t *testing.T) {
	poster := DefaultPoster(t)
	require.NotEqual(t, poster.DmensNftId, "")
}

func TestPostNote(t *testing.T) {
	poster := DefaultPoster(t)

	// chain := poster.chain
	chain := sui.NewChainWithRpcUrl("https://fullnode.testnet.sui.io")

	txn, err := poster.DmensPost("Hello world")
	require.Nil(t, err)

	fee, err := chain.EstimateGasFee(txn)
	require.Nil(t, err)
	t.Log(fee.Value)

	signed, err := txn.SignWithAccount(M1Account)
	require.Nil(t, err)

	hash, err := chain.SendRawTransaction(signed.Value)
	require.Nil(t, err)
	t.Log(hash)
}
