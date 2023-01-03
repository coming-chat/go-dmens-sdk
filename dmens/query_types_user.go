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

	array *base.AnyArray
}

func (n *UserPage) JsonString() (string, error) {
	return JsonString(n)
}

func (n *UserPage) FirstObject() *UserInfo {
	if len(n.Users) <= 0 {
		return nil
	}
	return &n.Users[0]
}

func (p *UserPage) UserArray() *base.AnyArray {
	if p.array == nil {
		a := make([]any, len(p.Users))
		for _, n := range p.Users {
			a = append(a, n)
		}
		p.array = &base.AnyArray{Values: a}
	}
	return p.array
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
