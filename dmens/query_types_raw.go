package dmens

import (
	"encoding/json"
	"time"
)

type rawFieldsId struct {
	Id struct {
		Id string `json:"id"`
	} `json:"id"`
}

type rawPageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
	// HasPreviousPage bool `json:"hasPreviousPage"`
	// StartCursor     string `json:"startCursor"`
}

type rawUser = UserInfo

type rawNote struct {
	CreateTime string `json:"createTime,omitempty"`
	ObjectId   string `json:"objectId,omitempty"`

	Fields *struct {
		// rawFieldsId
		// Name  string `json:"name"`
		Value struct {
			// Type string `json:"type"`
			Fields struct {
				Action NoteAction `json:"action"`
				Text   string     `json:"text"`
				Poster string     `json:"poster"`
				RefId  string     `json:"ref_id"`

				// rawFieldsId
				// Url    string     `json:"url"`
				// AppId  int        `json:"app_id"`
			} `json:"fields"`
		} `json:"value"`
	} `json:"fields,omitempty"`

	// Owner      interface{} `json:"owner,omitempty"`
	// UpdateTime string `json:"updateTime,omitempty"`
	// Status     string `json:"status,omitempty"`
	// DataType   string `json:"dataType,omitempty"`
	// Type       string `json:"type,omitempty"`
	// NodeId     string `json:"nodeId,omitempty"`
	// Digest     string `json:"digest,omitempty"`
	// Version    string `json:"version,omitempty"`
	// IsUpdate   bool   `json:"isUpdate,omitempty"`
	// StorageRebate       string `json:"storageRebate,omitempty"`
	// PreviousTransaction string `json:"previousTransaction,omitempty"`
	// HasPublicTransfer   bool   `json:"hasPublicTransfer,omitempty"`
}

func (a *rawNote) MapToNote() *Note {
	var timestamp int64 = 0
	t, err := time.Parse("2006-01-02T15:04:05.999999", a.CreateTime)
	if err == nil {
		timestamp = t.Unix()
	}
	fields := a.Fields.Value.Fields
	return &Note{
		CreateTime: timestamp,
		NoteId:     a.ObjectId,
		Action:     fields.Action,
		Text:       fields.Text,
		Poster:     fields.Poster,
		RefId:      fields.RefId,
	}
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

func (u *rawUserFollow) MapToUserInfo() *UserInfo {
	res := &UserInfo{Address: u.Fields.Name}
	_ = json.Unmarshal(u.Fields.Value, res)
	return res
}

type rawPageable[O rawUser | rawNote | rawUserFollow, M UserInfo | Note] struct {
	TotalCount int `json:"totalCount,omitempty"`
	Edges      []struct {
		Node   O      `json:"node"`
		Cursor string `json:"cursor"`
	} `json:"edges,omitempty"`
	PageInfo *rawPageInfo `json:"pageInfo,omitempty"`
}

func (p *rawPageable[O, M]) mapToSdkPage(pageSize int, maper func(O) *M) *sdkPageable[M] {
	currentCount := len(p.Edges)
	if currentCount == 0 {
		return &sdkPageable[M]{
			TotalCount_: p.TotalCount,
		}
	}
	items := make([]*M, currentCount)
	for idx, node := range p.Edges {
		items[idx] = maper(node.Node)
	}
	page := &sdkPageable[M]{
		TotalCount_:    p.TotalCount,
		CurrentCount_:  currentCount,
		CurrentCursor_: p.Edges[currentCount-1].Cursor,

		Items: items,
	}
	if p.PageInfo != nil {
		page.HasNextPage_ = p.PageInfo.HasNextPage
	} else {
		page.HasNextPage_ = currentCount >= pageSize
	}
	return page
}

func (p *rawPageable[O, M]) FirstObject() *O {
	if len(p.Edges) <= 0 {
		return nil
	}
	return &p.Edges[0].Node
}

type rawUserPage struct {
	rawPageable[rawUser, UserInfo]
}
type rawNotePage struct {
	rawPageable[rawNote, Note]
}
type rawUserFollowPage struct {
	rawPageable[rawUserFollow, UserInfo]
}
