package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func DefaultPoster(t *testing.T) *Poster {
	address := "0x6c5d2cd6e62734f61b4e318e58cbfd1c4b99dfaf"
	poster, err := NewPoster(&PosterConfig{Address: address}, DevnetConfig)
	require.Nil(t, err)
	return poster
}

func TestNewPoster(t *testing.T) {
	poster := DefaultPoster(t)
	require.Equal(t, poster.DmensNftId, "0xf9d0bc907daa0e8dbd93fcc44db4415e44d31d38")
}
