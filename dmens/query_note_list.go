package dmens

import (
	"fmt"
)

func (p *Poster) QueryReplyNoteList(noteId string, pageSize int, afterCursor string) (*NotePage, error) {
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, ref_id: "%v"}}}}`, ACTION_REPLY, noteId)
	return p.queryNoteList(pageSize, afterCursor, fieldJson)
}

// QueryUserNoteList
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserNoteList(user string, pageSize int, afterCursor string) (*NotePage, error) {
	if user == "" {
		user = p.Address
	}
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, poster: "%v"}}}}`, ACTION_POST, user)
	return p.queryNoteList(pageSize, afterCursor, fieldJson)
}

// QueryAllNoteList
// @param pageSize The number of notes per page.
// @param afterCursor Each page has a cursor, and you need to specify the cursor to get the next page of content, If you want to get the first page of content, pass in empty.
func (p *Poster) QueryAllNoteList(pageSize int, afterCursor string) (*NotePage, error) {
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v}}}}`, ACTION_POST)
	return p.queryNoteList(pageSize, afterCursor, fieldJson)
}

func (p *Poster) queryNoteList(pageSize int, afterCursor string, fieldJson string) (*NotePage, error) {
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

	return out.MapToNotePage(p), nil
}
