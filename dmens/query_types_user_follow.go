package dmens

import "encoding/json"

type UserFollowCount struct {
	User string `json:"user"`

	FollowerCount  int `json:"followerCount"`
	FollowingCount int `json:"followingCount"`
}

func (n *UserFollowCount) JsonString() (string, error) {
	return JsonString(n)
}

type rawUserFollow struct {
	// follower's params
	// FollowerAddress  string `json:"followerAddress"`  // we can use fields.name

	// following's params
	// FollowingAddress string `json:"followingAddress"` // we can use fields.name

	// Trend Users's params
	// FollowerNumber int    `json:"followerNumber,string"` // no need
	// Owner          string `json:"owner"` // we can use fields.name

	Fields struct {
		// rawFieldsId
		Name  string `json:"name"`
		Value []byte `json:"value"`
	} `json:"fields"`
}

type rawUserFollowPage struct {
	TotalCount int `json:"totalCount,omitempty"`
	Edges      []struct {
		Node   rawUserFollow `json:"node"`
		Cursor string        `json:"cursor"`
	} `json:"edges,omitempty"`
}

func (u *rawUserFollow) MapToUserInfo() *UserInfo {
	res := &UserInfo{Address: u.Fields.Name}
	json.Unmarshal(u.Fields.Value, res)
	return res
}

func (a *rawUserFollowPage) MapToUserPage() *UserPage {
	length := len(a.Edges)
	if length == 0 {
		return &UserPage{
			TotalCount: a.TotalCount,
		}
	}
	users := make([]UserInfo, 0)
	for _, n := range a.Edges {
		users = append(users, *n.Node.MapToUserInfo())
	}
	return &UserPage{
		TotalCount:    a.TotalCount,
		Users:         users,
		CurrentCursor: a.Edges[length-1].Cursor,
	}
}

func (a *rawUserFollowPage) FirstObject() *rawUserFollow {
	if len(a.Edges) <= 0 {
		return nil
	}
	return &a.Edges[0].Node
}
