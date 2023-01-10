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
}

func (u *UserInfo) JsonString() (string, error) {
	return JsonString(u)
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

func (n *UserFollowCount) JsonString() (string, error) {
	return JsonString(n)
}

type UserPage struct {
	*sdkPageable[UserInfo]
}
