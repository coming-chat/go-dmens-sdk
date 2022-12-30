package dmens

import (
	"fmt"

	"github.com/coming-chat/wallet-SDK/core/base"
)

type NoteStatus struct {
	NoteId string `json:"noteId"`
	Viewer string `json:"viewer"`

	ReplyCount  int `json:"replyCount"`
	RepostCount int `json:"repostCount"`
	LikeCount   int `json:"likeCount"`

	// Whether the viewer reposted it
	IsReposted bool `json:"isReposted"`
	// Whether the viewer liked it
	IsLiked bool `json:"isLiked"`
}

// @param noteId the note's id
// @param viewer the note's viewer, if the viewer is empty, the poster's address will be queried.
func (p *Poster) QueryNoteStatusById(noteId string, viewer string) (*NoteStatus, error) {
	var res = &NoteStatus{
		NoteId: noteId,
		Viewer: viewer,
	}
	if viewer == "" {
		viewer = p.Address
	}

	queryFormat := `
	query NoteStatus {
		allSuiObjects(
		  filter: {
			dataType: { equalTo: "moveObject" }
			status: { equalTo: "Exists" }
			type: { equalTo: "` + p.dmensObjectType() + `" }
			fields: { contains: {value: {fields: {ref_id: "` + noteId + `", action: %v %v }}}}
		  }
		) {
		  totalCount
		}
	  }
	`
	type action struct {
		queryString string
		action      func(status *NoteStatus, count int)
	}
	actions := map[string]action{
		"replyCount": {
			queryString: fmt.Sprintf(queryFormat, ACTION_REPLY, ""),
			action: func(status *NoteStatus, count int) {
				status.ReplyCount = count
			},
		},
		"repostCount": {
			queryString: fmt.Sprintf(queryFormat, ACTION_REPOST, ""),
			action: func(status *NoteStatus, count int) {
				status.RepostCount = count
			},
		},
		"isPosted": {
			queryString: fmt.Sprintf(queryFormat, ACTION_REPOST, `, poster: "`+viewer+`"`),
			action: func(status *NoteStatus, count int) {
				status.IsReposted = count > 0
			},
		},
		"likeCount": {
			queryString: fmt.Sprintf(queryFormat, ACTION_LIKE, ""),
			action: func(status *NoteStatus, count int) {
				status.LikeCount = count
			},
		},
		"isLiked": {
			queryString: fmt.Sprintf(queryFormat, ACTION_LIKE, `, poster: "`+viewer+`"`),
			action: func(status *NoteStatus, count int) {
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
