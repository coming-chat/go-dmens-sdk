package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func DefaultPoster(t *testing.T) *Poster {
	address := "0x6fc6148816617c3c3eccb1d09e930f73f6712c9c"
	poster, err := NewPoster(&PosterConfig{Address: address}, DevnetConfig)
	require.Nil(t, err)
	return poster
}

func TestNewPoster(t *testing.T) {
	poster := DefaultPoster(t)
	require.NotEqual(t, poster.DmensNftId, "")
}
