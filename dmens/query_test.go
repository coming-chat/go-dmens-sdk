package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryTwittersMyFollowed(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTwittersMyFollowed(10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryTrendTwitters(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendTwittersList(10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryTwitterByNoteId(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTwitterByNoteId("0xc4dbf29b5513f7695d6370e094c5ad03fb44acc2", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserTwitters(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserTwitters("", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserFollowing(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowing("")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserFollowers(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowers("0x3f432b985d6a5bd6f3b8f96a44f9adf272a59bb3")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryTrendUsers(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendUserList()
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserCharaters(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserCharacter("")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUsersByName(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUsersByName("g", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserInfoByAddress(t *testing.T) {
	poster := DefaultPoster(t)

	// query default poster's info
	res, err := poster.QueryUserInfoByAddress("")
	require.Nil(t, err)
	t.Log(res)

	// query specified user's info
	res, err = poster.QueryUserInfoByAddress("0x111")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserTwittersList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserTwittersList("", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryTwitterByRefId(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTwittersList(0, "", "", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}
