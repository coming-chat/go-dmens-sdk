package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchNftDetail(t *testing.T) {
	poster := DefaultPoster(t)
	nftId := "0xc9ef76e6f130dc70476601bbdf58242eb136942d" // invalid address
	nft, err := poster.QueryNFTAvatar(nftId)
	require.Error(t, err)
	t.Log(nft)
}

func TestQuerySuiNameByAddress(t *testing.T) {
	poster := DefaultPoster(t)

	owner := "0x5f2bd2399ec538a71f56b6928ca4d8b11200fe08" // invalid address
	name, err := poster.QuerySuiNameByAddress(owner)
	require.Error(t, err)

	owner = "0x57188743983628b3474648d8aa4a9ee8abebe8f6816243773d7e8ed4fd833a28"
	name, err = poster.QuerySuiNameByAddress(owner)
	require.NoError(t, err)
	t.Log(name.Value)
}
