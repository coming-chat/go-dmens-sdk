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

func TestBatchQueryNoteByIds(t *testing.T) {
	poster := DefaultPoster(t)
	ids := []string{
		"0xaea51c305a367ae8afdd788099c74dfdcf5d1a5b",
		"0xc61a46de4ca47f2ccc7f374c0161ab23d2391ada",
		"0xf64673b3cd979712915578a8e86ed79561e8578a",
	}
	res, err := poster.BatchQueryNoteByIds(ids)
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserFollowing(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowing("", 10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserFollowers(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowers("0x3f432b985d6a5bd6f3b8f96a44f9adf272a59bb3", 10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
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
	t.Log(res.JsonString())
}

func TestQueryUserCharaters(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.queryUserCharacter("")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUsersByName(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUsersByName("g", 10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserInfoByAddress(t *testing.T) {
	poster := DefaultPoster(t)

	// query default poster's info
	res, err := poster.QueryUserInfoByAddress("")
	require.Nil(t, err)
	t.Log(res.JsonString())

	// query specified user's info
	res, err = poster.QueryUserInfoByAddress("0x111")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestBatchQueryUserByAddressJson(t *testing.T) {
	poster := DefaultPoster(t)
	address := `[
		"0xbf16dbc4b99159f3afc8e5743de75e9c53c3f171",
		"0x7c1b34834f58064743252260eaefa9ce443b24ed",
		"0xfe443c8f33482b1d5165fbd8bc007c58bd1cab41"
	]`
	res, err := poster.BatchQueryUserByAddressJson(address)
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserNoteList("", 10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserRepostList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserRepostList("", 4, "")
	// res, err := poster.QueryUserRepostListAsNotePage("", 4, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryReplyNoteList(t *testing.T) {
	noteId := "0xd1b1fa2d807fac385b9e3897778091e6074942c4" // `123123@littlema @hi`
	noteId = "0xa9af508f0e489658905d7ae6d193864855d71e84"  // `😀😀😀`
	noteId = "0xc4dbf29b5513f7695d6370e094c5ad03fb44acc2"  // `@Chatgpt Are you ready?👻`

	poster := DefaultPoster(t)
	res, err := poster.QueryReplyNoteList(noteId, 10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryAllNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryAllNoteList(10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryNoteStatusById(t *testing.T) {
	noteId := "0xa8422cb686fea931a288fa63738ee0503b9ca53e"
	viewer := ""

	poster := DefaultPoster(t)
	res, err := poster.QueryNoteStatusById(noteId, viewer)
	require.Nil(t, err)
	t.Log(JsonString(res))
	t.Log("")
}

func TestBatchQueryNoteStatusByIds(t *testing.T) {
	noteids := []string{
		"0x57bcbc127d6ac3a26a5cf6bbfdefd04c2903740a",
		"0xf74ef6da596105f6596338ffc9b913a727237cc5",
	}
	poster := DefaultPoster(t)
	res, err := poster.BatchQueryNoteStatusByIds(noteids, "")
	require.Nil(t, err)
	t.Log(JsonString(res))
}

func TestIsMyFollowing(t *testing.T) {
	poster := DefaultPoster(t)
	isFollow, err := poster.IsMyFollowing("0x7c1b34834f58064743252260eaefa9ce443b24ed")
	require.Nil(t, err)
	t.Log(isFollow)
}

func TestAsUserInfo(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUsersByName("g", 10, "")
	require.Nil(t, err)

	userArray := res.ItemArray()
	for i := 0; i < res.CurrentCount(); i++ {
		userAny := userArray.ValueOf(i)
		user := AsUserInfo(userAny)
		t.Log(user.JsonString())
	}
}

func TestAsNote(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendNoteList(10, "")
	require.Nil(t, err)

	noteArray := res.ItemArray()
	for i := 0; i < res.CurrentCount(); i++ {
		noteAny := noteArray.ValueOf(i)
		note := AsNote(noteAny)
		t.Log(note.JsonString())
	}
}
