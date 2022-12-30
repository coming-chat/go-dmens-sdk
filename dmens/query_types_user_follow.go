package dmens

import "encoding/json"

type rawUserFollow struct {
	// FollowerAddress  string `json:"followerAddress"`  // we can use fields.name
	// FollowingAddress string `json:"followingAddress"` // we can use fields.name

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
