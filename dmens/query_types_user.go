package dmens

import (
	"encoding/json"

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

type UserPage struct {
	*Pageable
	Users []*UserInfo `json:"users"`
}

func (u *UserPage) JsonString() (string, error) {
	return JsonString(u)
}

func (u *UserPage) FirstObject() *UserInfo {
	if len(u.Users) <= 0 {
		return nil
	}
	return u.Users[0]
}

func (u *UserPage) UserArray() *base.AnyArray {
	if u.anyArray == nil {
		a := make([]any, len(u.Users))
		for idx, n := range u.Users {
			a[idx] = n
		}
		u.anyArray = &base.AnyArray{Values: a}
	}
	return u.anyArray
}

type UserFollowCount struct {
	User string `json:"user"`

	FollowerCount  int `json:"followerCount"`
	FollowingCount int `json:"followingCount"`
}

func (n *UserFollowCount) JsonString() (string, error) {
	return JsonString(n)
}

func (a *rawUserPage) MapToUserPage(pageSize int) *UserPage {
	users := make([]*UserInfo, len(a.Edges))
	for idx, n := range a.Edges {
		users[idx] = &n.Node
	}
	page := a.mapToBasePage(pageSize)
	return &UserPage{
		Pageable: page,
		Users:    users,
	}
}

func (u *rawUserFollow) MapToUserInfo() *UserInfo {
	res := &UserInfo{Address: u.Fields.Name}
	_ = json.Unmarshal(u.Fields.Value, res)
	return res
}

func (a *rawUserFollowPage) MapToUserPage(pageSize int) *UserPage {
	users := make([]*UserInfo, len(a.Edges))
	for idx, n := range a.Edges {
		users[idx] = n.Node.MapToUserInfo()
	}
	page := a.mapToBasePage(pageSize)
	return &UserPage{
		Pageable: page,
		Users:    users,
	}
}

func (a *rawUserFollowPage) FirstObject() *rawUserFollow {
	if len(a.Edges) <= 0 {
		return nil
	}
	return &a.Edges[0].Node
}
