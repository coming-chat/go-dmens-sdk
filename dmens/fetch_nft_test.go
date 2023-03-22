package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchNftDetail(t *testing.T) {
	poster := DefaultPoster(t)
	nftId := "0xc9ef76e6f130dc70476601bbdf58242eb136942d"
	nft, err := poster.QueryNFTAvatar(nftId)
	require.Nil(t, err)

	t.Log(nft)
}
