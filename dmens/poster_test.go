package dmens

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var whoami = ""
var M1 = os.Getenv("WalletSdkTestM1")

func init() {
	out, _ := exec.Command("whoami").Output()
	whoami = strings.TrimSpace(string(out))
}

func DefaultPoster(t *testing.T) *Poster {
	address := ""
	switch whoami {
	case "gg":
		address = "0x6c5d2cd6e62734f61b4e318e58cbfd1c4b99dfaf"
	default:
		address = "0x6fc6148816617c3c3eccb1d09e930f73f6712c9c"
	}
	poster, err := NewPoster(&PosterConfig{Address: address}, DevnetConfig)
	require.Nil(t, err)
	return poster
}

func TestNewPoster(t *testing.T) {
	poster := DefaultPoster(t)
	require.NotEqual(t, poster.DmensNftId, "")
}
