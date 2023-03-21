package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/base"
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

	IsFollowing bool `json:"isFollowing"`
}

func (u *UserInfo) JsonString() (*base.OptionalString, error) {
	return base.JsonString(u)
}
func NewUserInfoWithJsonString(str string) (*UserInfo, error) {
	var o UserInfo
	err := base.FromJsonString(str, &o)
	return &o, err
}

func (u *UserInfo) AsAny() *base.Any {
	return &base.Any{Value: u}
}

func AsUserInfo(a *base.Any) *UserInfo {
	if res, ok := a.Value.(*UserInfo); ok {
		return res
	}
	if res, ok := a.Value.(UserInfo); ok {
		return &res
	}
	return nil
}

type UserFollowCount struct {
	User string `json:"user"`

	FollowerCount  int `json:"followerCount"`
	FollowingCount int `json:"followingCount"`
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
	*sdkPageable[UserInfo]
}

func NewUserPageWithJsonString(str string) (*UserPage, error) {
	var o sdkPageable[UserInfo]
	err := base.FromJsonString(str, &o)
	if err != nil {
		return nil, err
	}
	return &UserPage{&o}, nil
}
