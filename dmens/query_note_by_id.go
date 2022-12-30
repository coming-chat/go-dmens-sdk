package dmens

func (p *Poster) QueryNoteById(noteId string) (string, error) {
	query := Query{
		Query: `
		query NoteById(
			$type: String
			$noteId: String
		  ) {
			allSuiObjects(
			  filter: {
				dataType: { equalTo: "moveObject" }
				status: { equalTo: "Exists" }
				type: { equalTo: $type }
				objectId: { equalTo: $noteId }
			  }
			  orderBy: CREATE_TIME_DESC
			  first: 10
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
			"type":   p.dmensObjectType(),
			"noteId": noteId,
		},
	}

	var out rawNotePage
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return "", err
	}
	return out.MapToNotePage().JsonString()
}
