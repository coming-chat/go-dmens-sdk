package dmens

func (p *Poster) QueryNotesMyFollowed(pageSize int, offset int) (string, error) {
	query := &Query{
		Query: `
		query NotesMyFollowed(
			$dmensMetaObjectType: String
			$dmensObjectType: String
			$objectOwner: String
			$first: Int
			$offset: Int
			$action: Int
		  ) {
			home(
			  dmensMetaObjectType: $dmensMetaObjectType
			  dmensObjectType: $dmensObjectType
			  objectOwner: $objectOwner
			  first: $first
			  offset: $offset
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
			"offset":              offset,
			"action":              ACTION_POST,
		},
	}

	var out rawNotePage
	err := p.makeQueryOut(query, "home", &out)
	if err != nil {
		return "", err
	}

	return out.MapToNotePage().JsonString()
}
