package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryNotesMyFollowed(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryNotesMyFollowed(10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryTrendNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendNoteList(10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryNoteById(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryNoteById("0xc4dbf29b5513f7695d6370e094c5ad03fb44acc2")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserNotes(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserNotes("", 10, 0)
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

func TestQueryUserNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserNoteList("", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryNoteList(0, "", "", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryNoteStatusById(t *testing.T) {
	noteId := "0xd1b1fa2d807fac385b9e3897778091e6074942c4"
	viewer := ""

	poster := DefaultPoster(t)
	res, err := poster.QueryNoteStatusById(noteId, viewer)
	require.Nil(t, err)
	t.Log(res)
}
