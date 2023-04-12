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
var M1Address = "0x7e875ea78ee09f08d72e2676cf84e0f1c8ac61d94fa339cc8e37cace85bebc6e"

func init() {
	out, _ := exec.Command("whoami").Output()
	whoami = strings.TrimSpace(string(out))
}

func DefaultPoster(t *testing.T) *Poster {
	address := ""
	switch whoami {
	case "gg":
		address = M1Address
	default:
		address = M1Address
	}
	poster, err := NewPoster(&PosterConfig{Address: address}, DevnetConfig)
	require.Nil(t, err)
	return poster
}

func TestNewPoster(t *testing.T) {
	poster := DefaultPoster(t)
	require.NotEqual(t, poster.DmensNftId, "")
}
