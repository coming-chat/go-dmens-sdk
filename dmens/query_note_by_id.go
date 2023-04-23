package dmens

import "fmt"

func (p *Poster) QueryNoteById(noteId string) (*Note, error) {
	page, err := p.BatchQueryNoteByIds([]string{noteId})
	if err != nil {
		return nil, err
	}
	if page.CurrentCount() <= 0 {
		return nil, nil
	}
	return page.Items[0], nil
}

func (p *Poster) BatchQueryNoteByIds(ids []string) (*NotePage, error) {
	if len(ids) <= 0 {
		return &NotePage{}, nil
	}
	filterAppid := ""
	if p.Reviewing {
		filterAppid = fmt.Sprintf("fields: { contains: {value: {fields: {app_id: %v}}}}", appIdForComingChatApp)
	}
	queryString := fmt.Sprintf(`
	query NoteById(
		$type: String
		$noteIds: [String!]
	  ) {
		allSuiObjects(
		  filter: {
			dataType: { equalTo: "moveObject" }
			status: { equalTo: "Exists" }
			type: { equalTo: $type }
			objectId: { in: $noteIds } %v
		  }
		) {
		  totalCount
		  edges {
			cursor
			node {
			  createTime
			  fields
			  objectId
			}
		  }
		}
	  }
	`, filterAppid)
	query := Query{
		Query: queryString,
		Variables: map[string]interface{}{
			"type":    p.dmensObjectType(),
			"noteIds": ids,
		},
	}

	var out rawNotePage
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return nil, err
	}

	page := out.MapToNotePage(p, len(ids))
	return page, nil
}
