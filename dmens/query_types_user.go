package dmens

import "github.com/coming-chat/wallet-SDK/core/base"

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
	return nil
}

type UserPage struct {
	Users         []UserInfo `json:"users"`
	CurrentCursor string     `json:"currentCursor"`
	CurrentCount  int        `json:"currentCount"`
	TotalCount    int        `json:"totalCount"`

	usersArray *base.AnyArray
}

func (u *UserPage) JsonString() (string, error) {
	return JsonString(u)
}

func (u *UserPage) FirstObject() *UserInfo {
	if len(u.Users) <= 0 {
		return nil
	}
	return &u.Users[0]
}

func (u *UserPage) UserArray() *base.AnyArray {
	if u.usersArray == nil {
		a := make([]any, len(u.Users))
		for _, n := range u.Users {
			a = append(a, n)
		}
		u.usersArray = &base.AnyArray{Values: a}
	}
	return u.usersArray
}

type rawUserPage struct {
	TotalCount int `json:"totalCount,omitempty"`
	Edges      []struct {
		Node   UserInfo `json:"node"`
		Cursor string   `json:"cursor"`
	} `json:"edges,omitempty"`
}

func (a *rawUserPage) MapToUserPage() *UserPage {
	length := len(a.Edges)
	if length == 0 {
		return &UserPage{
			TotalCount: a.TotalCount,
		}
	}
	users := make([]UserInfo, 0)
	for _, n := range a.Edges {
		users = append(users, n.Node)
	}
	return &UserPage{
		TotalCount:    a.TotalCount,
		Users:         users,
		CurrentCount:  len(users),
		CurrentCursor: a.Edges[length-1].Cursor,
	}
}
