package dmens

func (p *Poster) QueryNotesMyFollowed(pageSize int, afterCursor string) (*NotePage, error) {
	query := &Query{
		Query: `
		query NotesMyFollowed(
			$dmensMetaObjectType: String
			$dmensObjectType: String
			$objectOwner: String
			$first: Int
			$action: String
		  ) {
			home(
			  dmensMetaObjectType: $dmensMetaObjectType
			  dmensObjectType: $dmensObjectType
			  objectOwner: $objectOwner
			  first: $first
			  after: #cursor#
			  filter: {
				status: { equalTo: "Exists" }
				fields: { contains: { value: { fields: { action: $action } } } }
			  }
			) {
			  totalCount
			  edges {
				cursor
				node {
				  createTime
				  objectId
				  fields
				}
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"dmensMetaObjectType": p.dmensMetaObjectType(),
			"dmensObjectType":     p.dmensObjectType(),
			"objectOwner":         p.Address,
			"first":               pageSize,
			"action":              "0", //ACTION_POST,
		},
		Cursor: afterCursor,
	}

	var out rawNotePage
	err := p.makeQueryOut(query, "home", &out)
	if err != nil {
		return nil, err
	}

	return out.MapToNotePage(p, pageSize), nil
}
