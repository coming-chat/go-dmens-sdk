package dmens

import "github.com/coming-chat/wallet-SDK/core/base"

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

type rawPageable[T rawUser | rawNote | rawUserFollow] struct {
	TotalCount int `json:"totalCount,omitempty"`
	Edges      []struct {
		Node   T      `json:"node"`
		Cursor string `json:"cursor"`
	} `json:"edges,omitempty"`
	PageInfo *rawPageInfo `json:"pageInfo,omitempty"`
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

type rawUserPage struct {
	rawPageable[rawUser]
}
type rawNotePage struct {
	rawPageable[rawNote]
}
type rawUserFollowPage struct {
	rawPageable[rawUserFollow]
}

func (p *rawPageable[T]) mapToBasePage(pageSize int) *Pageable {
	currentCount := len(p.Edges)
	if currentCount == 0 {
		return &Pageable{
			TotalCount: p.TotalCount,
		}
	}
	page := &Pageable{
		TotalCount:    p.TotalCount,
		CurrentCount:  currentCount,
		CurrentCursor: p.Edges[currentCount-1].Cursor,
	}
	if p.PageInfo != nil {
		page.HasNextPage = p.PageInfo.HasNextPage
	} else {
		page.HasNextPage = currentCount >= pageSize
	}
	return page
}

type Pageable struct {
	TotalCount    int    `json:"totalCount"`
	CurrentCount  int    `json:"currentCount"`
	CurrentCursor string `json:"currentCursor"`
	HasNextPage   bool   `json:"hasNextPage"`

	anyArray *base.AnyArray
}
