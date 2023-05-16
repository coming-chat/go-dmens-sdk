package dmens

import (
	"testing"

	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/stretchr/testify/require"
)

func TestFollow(t *testing.T) {
	poster := DefaultPoster(t)
	chain := poster.chain

	address := "0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5"
	array := base.StringArray{Values: []string{address}}

	txn, err := poster.DmensFollow(&array)
	require.Nil(t, err)

	simulateCheck(t, chain, txn, true)

	// acc, err := sui.NewAccountWithMnemonic(M1)
	// require.Nil(t, err)
	// signedTxn, err := txn.SignWithAccount(acc)
	// require.Nil(t, err)
	// hash, err := poster.chain.SendRawTransaction(signedTxn.Value)
	// require.Nil(t, err)
	// t.Log(hash)
}

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
	res, err := poster.QueryNoteById("0x53b97587007e3a4e51deed1de8e1ebfbc3a6a2f6a0ce6a97836e87842b9920ed")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestBatchQueryNoteByIds(t *testing.T) {
	poster := DefaultPoster(t)
	ids := []string{
		"0x53b97587007e3a4e51deed1de8e1ebfbc3a6a2f6a0ce6a97836e87842b9920ed",
		"0xc61a46de4ca47f2ccc7f374c0161ab23d2391ada",
	}
	res, err := poster.BatchQueryNoteByIds(ids)
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserFollowing(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowing("0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5", 5, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserFollowers(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserFollowers("0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5", 5, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserFollowCount(t *testing.T) {
	poster := DefaultPoster(t)

	res, err := poster.QueryUserFollowCount("")
	require.Nil(t, err)
	t.Log(res.JsonString())

	res, err = poster.QueryUserFollowCount("0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryTrendUsers(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendUserList(10)
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
	// res, err = poster.QueryUserInfoByAddress("0x5cbf57ec2dd5c4eb0560ee6ac20e9f8f3a75eca1")

	// query user info have nft avatar
	res, err = poster.QueryUserInfoByAddress("0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestBatchQueryUserByAddressJson(t *testing.T) {
	poster := DefaultPoster(t)
	address := `[
		"0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5",
		"0x7c1b34834f58064743252260eaefa9ce443b24ed",
		"0xfe443c8f33482b1d5165fbd8bc007c58bd1cab41"
	]`
	res, err := poster.BatchQueryUserByAddressJson(address)
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserNoteList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserNoteList("0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5", 10, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryUserRepostList(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryUserRepostList("0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5", 4, "")
	// res, err := poster.QueryUserRepostListAsNotePage("", 4, "")
	require.Nil(t, err)
	t.Log(res.JsonString())
}

func TestQueryReplyNoteList(t *testing.T) {
	noteId := "0x53b97587007e3a4e51deed1de8e1ebfbc3a6a2f6a0ce6a97836e87842b9920ed"

	poster := DefaultPoster(t)
	res, err := poster.QueryReplyNoteList(noteId, 7, "")
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
	noteId := "0x53b97587007e3a4e51deed1de8e1ebfbc3a6a2f6a0ce6a97836e87842b9920ed"
	viewer := ""

	poster := DefaultPoster(t)
	res, err := poster.QueryNoteStatusById(noteId, viewer)
	require.Nil(t, err)
	t.Log(base.JsonString(res))
	t.Log("")
}

func TestBatchQueryNoteStatusByIds(t *testing.T) {
	noteids := []string{
		"0x53b97587007e3a4e51deed1de8e1ebfbc3a6a2f6a0ce6a97836e87842b9920ed",
		"0xa3285d323d8f36bed42641f7ae80e005c3ce290de00d3cdf98fe2e79efa288da",
		"0xf74ef6da596105f6596338ffc9b913a727237cc5",
	}
	poster := DefaultPoster(t)
	res, err := poster.BatchQueryNoteStatusByIds(noteids, "")
	require.Nil(t, err)
	t.Log(base.JsonString(res))
}

func TestIsMyFollowing(t *testing.T) {
	poster := DefaultPoster(t)
	isFollow, err := poster.IsMyFollowing("0x919da05ad6103df99c988912a2d39475fcc672e2666540ce02e049649957d0b5")
	require.Nil(t, err)
	t.Log(isFollow)
}

func TestBatchQueryFollowingStatus(t *testing.T) {
	poster := DefaultPoster(t)
	viewer := "0x6d72b1de53d114352b9996a6c1a573a142f284e4"
	users := []string{
		"0x6f237da16dc1fede1bad6a250b03137cd8d9aef8", // true
		"0x24cf35e631c4a5006789f0575a6b470160b887b5", // true
		"0x7bc358b4e2e57332cb38a448740219be360b60ea", // false
		"0x32592deb1f071d451a2d93fe34851e2229cd6635", // true
		"0x6d72b1de53d114352b9996a6c1a573a142f284e4", // false
	}
	status, err := poster.batchQueryIsFollowingStatus(viewer, users)
	require.Nil(t, err)
	t.Log(status)
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
	res, err := poster.QueryNotesMyFollowed(10, "")
	// res, err := poster.QueryTrendNoteList(10, "")
	require.Nil(t, err)

	noteArray := res.ItemArray()
	for i := 0; i < res.CurrentCount(); i++ {
		noteAny := noteArray.ValueOf(i)
		note := AsNote(noteAny)
		t.Log(note.JsonString())
	}
}

func TestJsonable(t *testing.T) {
	poster := DefaultPoster(t)

	// --------------- NotePage
	notes, err := poster.QueryTrendNoteList(3, "")
	require.Nil(t, err)
	require.GreaterOrEqual(t, notes.TotalCount(), 1)
	require.LessOrEqual(t, notes.CurrentCount(), 3)

	jsonString, err := notes.JsonString()
	require.Nil(t, err)

	newNotes, err := NewNotePageWithJsonString(jsonString.Value)
	require.Nil(t, err)
	require.Equal(t, notes, newNotes)

	// --------------- UserPage
	users, err := poster.QueryTrendUserList(3)
	require.Nil(t, err)
	require.GreaterOrEqual(t, users.TotalCount(), 1)
	require.LessOrEqual(t, users.CurrentCount(), 3)

	jsonString, err = users.JsonString()
	require.Nil(t, err)

	newUsers, err := NewUserPageWithJsonString(jsonString.Value)
	require.Nil(t, err)
	require.Equal(t, users, newUsers)
}
