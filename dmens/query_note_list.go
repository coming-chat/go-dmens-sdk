package dmens

import (
	"fmt"
)

func (p *Poster) QueryReplyNoteList(noteId string, pageSize, offset int) (string, error) {
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, ref_id: "%v"}}}}`, ACTION_REPLY, noteId)
	return p.queryNoteList(pageSize, offset, fieldJson)
}

// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserNoteList(user string, pageSize, offset int) (string, error) {
	if user == "" {
		user = p.Address
	}
	fieldJson := fmt.Sprintf(`fields: { contains: {value: {fields: {action: %v, poster: "%v"}}}}`, ACTION_POST, user)
	return p.queryNoteList(pageSize, offset, fieldJson)
}

func (p *Poster) queryNoteList(pageSize, offset int, fieldJson string) (string, error) {
	queryString := fmt.Sprintf(`
	query NoteLists($type: String, $first: Int, $offset: Int) {
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
			"offset":    offset,
			"fieldJson": fieldJson,
		},
	}

	var out rawNotePage
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return "", err
	}

	return out.MapToNotePage().JsonString()
}
