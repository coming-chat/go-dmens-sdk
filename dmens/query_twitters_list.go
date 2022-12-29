package dmens

import (
	"encoding/json"
	"fmt"

	"github.com/coming-chat/wallet-SDK/core/base"
)

type TwitterStatus struct {
	TwitterId string `json:"twitterId"`
	Viewer    string `json:"viewer"`

	ReplyCount  int `json:"replyCount"`
	RepostCount int `json:"repostCount"`
	LikeCount   int `json:"likeCount"`

	// Whether the viewer reposted it
	IsReposted bool `json:"isReposted"`
	// Whether the viewer liked it
	IsLiked bool `json:"isLiked"`
}

// @param twitterId the twitter's id
// @param viewer the twitter's viewer, if the viewer is empty, the poster's address will be queried.
func (p *Poster) QueryTwitterStatusById(twitterId string, viewer string) (*TwitterStatus, error) {
	var res = &TwitterStatus{
		TwitterId: twitterId,
		Viewer:    viewer,
	}
	if viewer == "" {
		viewer = p.Address
	}

	queryFormat := `
	query TwittersStatus {
		allSuiObjects(
		  filter: {
			dataType: { equalTo: "moveObject" }
			status: { equalTo: "Exists" }
			type: { equalTo: "` + p.dmensObjectType() + `" }
			fields: { contains: {value: {fields: {ref_id: "` + twitterId + `", action: %v %v }}}}
		  }
		) {
		  totalCount
		}
	  }
	`
	type action struct {
		queryString string
		action      func(status *TwitterStatus, count int)
	}
	actions := map[string]action{
		"replyCount": {
			queryString: fmt.Sprintf(queryFormat, ACTION_REPLY, ""),
			action: func(status *TwitterStatus, count int) {
				status.ReplyCount = count
			},
		},
		"repostCount": {
			queryString: fmt.Sprintf(queryFormat, ACTION_REPOST, ""),
			action: func(status *TwitterStatus, count int) {
				status.RepostCount = count
			},
		},
		"isPosted": {
			queryString: fmt.Sprintf(queryFormat, ACTION_REPOST, `, poster: "`+viewer+`"`),
			action: func(status *TwitterStatus, count int) {
				status.IsReposted = count > 0
			},
		},
		"likeCount": {
			queryString: fmt.Sprintf(queryFormat, ACTION_LIKE, ""),
			action: func(status *TwitterStatus, count int) {
				status.LikeCount = count
			},
		},
		"isLiked": {
			queryString: fmt.Sprintf(queryFormat, ACTION_LIKE, `, poster: "`+viewer+`"`),
			action: func(status *TwitterStatus, count int) {
				status.IsLiked = count > 0
			},
		},
	}
	actionKeys := []string{"replyCount", "repostCount", "likeCount", "isPosted", "isLiked"}

	_, err := base.MapListConcurrentStringToString(actionKeys, func(key string) (string, error) {
		action := actions[key]

		query := Query{Query: action.queryString}
		var count int
		err := p.makeQueryOut(&query, "allSuiObjects.totalCount", &count)
		if err != nil {
			return "", err
		}
		action.action(res, count)
		return "", nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Poster) QueryTwittersList(action int, refId string, poster string, pageSize, offset int) (string, error) {
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, ref_id: "%v", poster:"%v"}}}}`, action, refId, poster)
	return p.queryTwittersList(pageSize, offset, fieldJson)
}

// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserTwittersList(user string, pageSize, offset int) (string, error) {
	if user == "" {
		user = p.Address
	}
	fieldJson := `fields: { contains: {value: {fields: {action: 0, poster: "` + user + `"}}}}`
	return p.queryTwittersList(pageSize, offset, fieldJson)
}

func (p *Poster) queryTwittersList(pageSize, offset int, fieldJson string) (string, error) {
	queryString := fmt.Sprintf(`
	query TwittersLists(
		$type: String
		$first: Int
		$offset: Int
	  ) {
		allSuiObjects(
		  filter: {
			dataType: { equalTo: "moveObject" }
			status: { equalTo: "Exists" }
			type: { equalTo: $type }
			%v
		  }
		  orderBy: CREATE_TIME_DESC
		  first: $first
		  offset: $offset
		) {
		  totalCount
		  pageInfo {
			hasNextPage
		  }
		  nodes {
			createTime
			dataType
			fields
			digest
			hasPublicTransfer
			isUpdate
			nodeId
			objectId
			owner
			previousTransaction
			status
			storageRebate
			type
			updateTime
			version
		  }
		}
	  }
	`, fieldJson)

	query := Query{
		Query: queryString,
		Variables: map[string]interface{}{
			"type":      p.dmensObjectType(),
			"first":     pageSize,
			"offset":    offset,
			"fieldJson": fieldJson,
		},
	}

	var out json.RawMessage
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
