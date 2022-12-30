package dmens

type User struct {
	FollowerNumber int    `json:"followerNumber,string"`
	Owner          string `json:"owner"`
}

type UserInfo struct {
	Address string `json:"address"`
	Avatar  string `json:"avatar"`
	Bio     string `json:"bio"`
	Name    string `json:"name"`
	NodeId  string `json:"nodeId"`
}

type UserPage struct {
	Users         []UserInfo `json:"users"`
	CurrentCursor string     `json:"currentCursor"`
	TotalCount    int        `json:"totalCount"`
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
		CurrentCursor: a.Edges[length-1].Cursor,
	}
}

func (n *UserInfo) JsonString() (string, error) {
	return JsonString(n)
}

func (n *UserPage) JsonString() (string, error) {
	return JsonString(n)
}
