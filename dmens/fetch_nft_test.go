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

func TestQuerySuiNameByAddress(t *testing.T) {
	poster := DefaultPoster(t)
	owner := "0x5f2bd2399ec538a71f56b6928ca4d8b11200fe08"
	name, err := poster.QuerySuiNameByAddress(owner)
	require.Nil(t, err)
	t.Log(name.Value)
}
