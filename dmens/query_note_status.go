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
		execute     func(status *NoteStatus, count int)
	}
	actions := []interface{}{
		&action{
			queryString: fmt.Sprintf(queryFormat, ACTION_REPLY, ""),
			execute: func(status *NoteStatus, count int) {
				status.ReplyCount = count
			},
		},
		&action{
			queryString: fmt.Sprintf(queryFormat, ACTION_REPOST, ""),
			execute: func(status *NoteStatus, count int) {
				status.RepostCount = count
			},
		},
		&action{
			queryString: fmt.Sprintf(queryFormat, ACTION_REPOST, `, poster: "`+viewer+`"`),
			execute: func(status *NoteStatus, count int) {
				status.IsReposted = count > 0
			},
		},
		&action{
			queryString: fmt.Sprintf(queryFormat, ACTION_LIKE, ""),
			execute: func(status *NoteStatus, count int) {
				status.LikeCount = count
			},
		},
		&action{
			queryString: fmt.Sprintf(queryFormat, ACTION_LIKE, `, poster: "`+viewer+`"`),
			execute: func(status *NoteStatus, count int) {
				status.IsLiked = count > 0
			},
		},
	}

	_, err := base.MapListConcurrent(actions, 5, func(i interface{}) (interface{}, error) {
		action, ok := i.(*action)
		if !ok {
			return i, nil
		}
		query := Query{Query: action.queryString}
		var count int
		err := p.makeQueryOut(&query, "allSuiObjects.totalCount", &count)
		if err != nil {
			return nil, err
		}
		action.execute(res, count)
		return i, nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 批量查询 page 中所有 note 的状态，数据会直接同步到 page 中每一个 note 对象中
// @param viewer the note's viewer, if the viewer is empty, the poster's address will be queried.
func (p *Poster) BatchQueryNoteStatus(page *NotePage, viewer string) error {
	if len(page.Notes) == 0 {
		return nil
	}
	if viewer == "" {
		viewer = p.Address
	}
	notesList := make([]interface{}, len(page.Notes))
	for idx, n := range page.Notes {
		notesList[idx] = n
	}
	_, err := base.MapListConcurrent(notesList, 5, func(i interface{}) (interface{}, error) {
		note, ok := i.(*Note)
		if !ok {
			return i, nil
		}
		status, err := p.QueryNoteStatusById(note.NoteId, viewer)
		if err != nil {
			return nil, err
		}
		note.Status = status
		return i, nil
	})
	return err
}
