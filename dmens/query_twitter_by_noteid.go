package dmens

import "encoding/json"

func (p *Poster) QueryTwitterByNoteId(noteId string, pageSize, offset int) (string, error) {
	query := Query{
		Query: `
		query TwitterByNoteId(
			$type: String
			$noteId: String
			$first: Int = 10
			$offset: Int = 0
		  ) {
			allSuiObjects(
			  filter: {
				dataType: { equalTo: "moveObject" }
				status: { equalTo: "Exists" }
				type: { equalTo: $type }
				objectId: { equalTo: $noteId }
			  }
			  orderBy: CREATE_TIME_DESC
			  first: $first
			  offset: $offset
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
			"type":   p.dmensObjectType(),
			"noteId": noteId,
			"first":  pageSize,
			"offset": offset,
		},
	}

	var out struct {
		TotalCount int `json:"totalCount"`
		PageInfo   struct {
			HasNextPage bool `json:"hasNextPage"`
		} `json:"pageInfo"`
		Nodes json.RawMessage `json:"nodes"`
	}
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return "", err
	}
	var res = map[string]interface{}{
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
