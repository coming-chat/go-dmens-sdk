package dmens

func (p *Poster) QueryNoteById(noteId string) (*Note, error) {
	page, err := p.BatchQueryNoteByIds([]string{noteId})
	if err != nil {
		return nil, err
	}
	if page.CurrentCount <= 0 {
		return nil, nil
	}
	return page.Notes[0], nil
}

func (p *Poster) BatchQueryNoteByIds(ids []string) (*NotePage, error) {
	if len(ids) <= 0 {
		return &NotePage{}, nil
	}
	query := Query{
		Query: `
		query NoteById(
			$type: String
			$noteIds: [String!]
		  ) {
			allSuiObjects(
			  filter: {
				dataType: { equalTo: "moveObject" }
				status: { equalTo: "Exists" }
				type: { equalTo: $type }
				objectId: { in: $noteIds }
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
		`,
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
