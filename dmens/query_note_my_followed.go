package dmens

import "encoding/json"

func (p *Poster) QueryNotesMyFollowed(pageSize int, offset int) (string, error) {
	query := &Query{
		Query: `
		query NotesMyFollowed(
			$dmensMetaObjectType: String
			$dmensObjectType: String
			$objectOwner: String
			$first: Int
			$offset: Int
			$fieldJson: JSON
		  ) {
			home(
			  dmensMetaObjectType: $dmensMetaObjectType
			  dmensObjectType: $dmensObjectType
			  objectOwner: $objectOwner
			  first: $first
			  offset: $offset
			  filter: { status: { equalTo: "Exists" }, fields: { contains: $fieldJson } }
			) {
			  totalCount
			  pageInfo {
				hasNextPage
			  }
			  nodes {
				createTime
				dataType
				digest
				fields
				hasPublicTransfer
				isUpdate
				nodeId
				objectId
				owner
				previousTransaction
				status
				storageRebate
				type
				updateTime
				version
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"dmensMetaObjectType": p.dmensMetaObjectType(),
			"dmensObjectType":     p.dmensObjectType(),
			"objectOwner":         p.followsId,
			"first":               pageSize,
			"offset":              offset,
			"fieldJson": map[string]map[string]map[string]NoteAction{
				"value": {"fields": {"action": ACTION_POST}},
			},
		},
	}

	var out struct {
		TotalCount int `json:"totalCount"`
		PageInfo   struct {
			HasNextPage bool `json:"hasNextPage"`
		} `json:"pageInfo"`
		Nodes json.RawMessage `json:"nodes"`
	}
	err := p.makeQueryOut(query, "home", &out)
	if err != nil {
		return "", err
	}

	res := map[string]interface{}{
		"totalCount":  out.TotalCount,
		"hasNextPage": out.PageInfo.HasNextPage,
		"nodes":       out.Nodes,
	}
	data, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
