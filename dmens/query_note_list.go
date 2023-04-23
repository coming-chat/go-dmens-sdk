package dmens

import (
	"fmt"
)

func (p *Poster) QueryReplyNoteList(noteId string, pageSize int, afterCursor string) (*NotePage, error) {
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, ref_id: "%v"}}}}`, ACTION_REPLY, noteId)
	return p.queryNoteList(pageSize, afterCursor, fieldJson, true)
}

// QueryUserNoteList
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserNoteList(user string, pageSize int, afterCursor string) (*NotePage, error) {
	if user == "" {
		user = p.Address
	}
	if p.Reviewing {
		fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, poster: "%v", app_id: %v}}}}`, ACTION_POST, user, appIdForComingChatApp)
		return p.queryNoteList(pageSize, afterCursor, fieldJson, true)
	} else {
		fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, poster: "%v"}}}}`, ACTION_POST, user)
		return p.queryNoteList(pageSize, afterCursor, fieldJson, true)
	}
}

// QueryUserRepostList
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserRepostList(user string, pageSize int, afterCursor string) (*RepostNotePage, error) {
	repostPage, originPage, err := p.queryUserRepostList(user, pageSize, afterCursor)
	if err != nil {
		return nil, err
	}
	page := combineRepostPage(repostPage, originPage)
	return page, nil
}

// QueryUserRepostListAsNotePage
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserRepostListAsNotePage(user string, pageSize int, afterCursor string) (*NotePage, error) {
	repostPage, originPage, err := p.queryUserRepostList(user, pageSize, afterCursor)
	if err != nil {
		return nil, err
	}
	page := combineRepostPage(repostPage, originPage)

	originNotes := make([]*Note, 0)
	for _, repost := range page.Items {
		originNotes = append(originNotes, repost.Note)
	}
	resPage := &NotePage{
		sdkPageable: &sdkPageable[Note]{
			Items:          originNotes,
			CurrentCount_:  len(originNotes),
			TotalCount_:    repostPage.TotalCount_,
			CurrentCursor_: repostPage.CurrentCursor_,
			HasNextPage_:   repostPage.HasNextPage_,
		},
	}
	return resPage, nil
}

func (p *Poster) queryUserRepostList(user string, pageSize int, afterCursor string) (*NotePage, *NotePage, error) {
	if user == "" {
		user = p.Address
	}
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, poster: "%v"}}}}`, ACTION_REPOST, user)
	repostPage, err := p.queryNoteList(pageSize, afterCursor, fieldJson, false)
	if err != nil {
		return nil, nil, err
	}
	if repostPage.CurrentCount() == 0 {
		return repostPage, &NotePage{sdkPageable: &sdkPageable[Note]{}}, nil
	}

	originNoteIds := make([]string, len(repostPage.Items))
	for idx, note := range repostPage.Items {
		originNoteIds[idx] = note.RefId
	}
	originNotePage, err := p.BatchQueryNoteByIds(originNoteIds)
	if err != nil {
		return nil, nil, err
	}
	return repostPage, originNotePage, nil
}

// QueryAllNoteList
// @param pageSize The number of notes per page.
// @param afterCursor Each page has a cursor, and you need to specify the cursor to get the next page of content, If you want to get the first page of content, pass in empty.
func (p *Poster) QueryAllNoteList(pageSize int, afterCursor string) (*NotePage, error) {
	if p.Reviewing {
		fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, app_id: %v}}}}`, ACTION_POST, appIdForComingChatApp)
		return p.queryNoteList(pageSize, afterCursor, fieldJson, true)
	} else {
		fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v}}}}`, ACTION_POST)
		return p.queryNoteList(pageSize, afterCursor, fieldJson, true)
	}
}

func (p *Poster) queryNoteList(pageSize int, afterCursor string, fieldJson string, needStatus bool) (*NotePage, error) {
	queryString := fmt.Sprintf(`
	query NoteLists($type: String, $first: Int) {
		allSuiObjects(
		  filter: {
			dataType: { equalTo: "moveObject" }
			status: { equalTo: "Exists" }
			type: { equalTo: $type }
			%v
		  }
		  orderBy: CREATE_TIME_DESC
		  first: $first
		  after: #cursor#
		) {
		  totalCount
		  edges {
			cursor
			node {
			  objectId
			  fields
			  createTime
			}
		  }
		}
	  }
	`, fieldJson)

	query := Query{
		Query: queryString,
		Variables: map[string]interface{}{
			"type":      p.dmensObjectType(),
			"first":     pageSize,
			"fieldJson": fieldJson,
		},
		Cursor: afterCursor,
	}

	var out rawNotePage
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return nil, err
	}

	if needStatus {
		return out.MapToNotePage(p, pageSize), nil
	} else {
		return out.MapToNotePage(nil, pageSize), nil
	}
}
