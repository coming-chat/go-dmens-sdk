package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/core/base/inter"
)

type UserInfo struct {
	Address string `json:"address"`
	Avatar  string `json:"avatar"`
	Bio     string `json:"bio"`
	Name    string `json:"name"`
	NodeId  string `json:"nodeId"`

	Background     string `json:"background"`
	Website        string `json:"website"`
	Identification string `json:"identification"`

	// Only queried when call QueryUserInfoByAddress
	SuiName string `json:"suiName"`

	// NFT avatar's nftid
	Item      []string   `json:"item"`
	NFTAvatar *NFTAvatar `json:"nftAvatar"`

	IsFollowing bool `json:"isFollowing"`
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func (u *UserInfo) JsonString() (*base.OptionalString, error) {
	return base.JsonString(u)
}
func NewUserInfoWithJsonString(str string) (*UserInfo, error) {
	var o UserInfo
	err := base.FromJsonString(str, &o)
	return &o, err
}

type UserFollowCount struct {
	User string `json:"user"`

	FollowerCount  int `json:"followerCount"`
	FollowingCount int `json:"followingCount"`
}

func NewUserFollowCount() *UserFollowCount {
	return &UserFollowCount{}
}

func (u *UserFollowCount) JsonString() (*base.OptionalString, error) {
	return base.JsonString(u)
}
func NewUserFollowCountWithJsonString(str string) (*UserFollowCount, error) {
	var o UserFollowCount
	err := base.FromJsonString(str, &o)
	return &o, err
}

type UserPage struct {
	*inter.SdkPageable[*UserInfo]
}

func NewUserPage() *UserPage {
	return &UserPage{}
}

func NewUserPageWithJsonString(str string) (*UserPage, error) {
	var o inter.SdkPageable[*UserInfo]
	err := base.FromJsonString(str, &o)
	if err != nil {
		return nil, err
	}
	return &UserPage{&o}, nil
}
