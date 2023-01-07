package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryNotesMyFollowed(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryNotesMyFollowed(10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryTrendNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendNoteList(10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryNoteById(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryNoteById("0x9013af3c543853805ce1532b76aac5ea1e25b368")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserFollowing(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowing("", 10, "")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserFollowers(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowers("0x3f432b985d6a5bd6f3b8f96a44f9adf272a59bb3", 10, "")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserFollowCount(t *testing.T) {
	poster := DefaultPoster(t)

	res, err := poster.QueryUserFollowCount("")
	require.Nil(t, err)
	t.Log(res.JsonString())

	res, err = poster.QueryUserFollowCount("0x3f432b985d6a5bd6f3b8f96a44f9adf272a59bb3")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryTrendUsers(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendUserList(10)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUserCharaters(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.queryUserCharacter("")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryUsersByName(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUsersByName("g", 10, "")
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
	res, err := poster.QueryUserNoteList("", 10, "")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryReplyNoteList(t *testing.T) {
	noteId := "0xd1b1fa2d807fac385b9e3897778091e6074942c4" // `123123@littlema @hi`
	noteId = "0xa9af508f0e489658905d7ae6d193864855d71e84"  // `😀😀😀`
	noteId = "0xc4dbf29b5513f7695d6370e094c5ad03fb44acc2"  // `@Chatgpt Are you ready?👻`

	poster := DefaultPoster(t)
	res, err := poster.QueryReplyNoteList(noteId, 10, "")
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryAllNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryAllNoteList(10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryNoteStatusById(t *testing.T) {
	noteId := "0xd1b1fa2d807fac385b9e3897778091e6074942c4"
	viewer := ""

	poster := DefaultPoster(t)
	res, err := poster.QueryNoteStatusById(noteId, viewer)
	require.Nil(t, err)
	t.Log(res)
}

func TestIsMyFollowing(t *testing.T) {
	poster := DefaultPoster(t)
	isFollow, err := poster.IsMyFollowing("0x7c1b34834f58064743252260eaefa9ce443b24ed")
	require.Nil(t, err)
	t.Log(isFollow)
}
