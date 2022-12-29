package dmens

import "encoding/json"

// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserTwitters(user string, pageSize, offset int) (string, error) {
	if user == "" {
		user = p.Address
	}
	query := Query{
		Query: `
		query userTwitters($type: String, $fieldJson: JSON, $first: Int, $offset: Int) {
			allSuiObjects(
			  filter: {
				status: { equalTo: "Exists" }
				dataType: { equalTo: "moveObject" }
				type: { equalTo: $type }
				fields: { contains: $fieldJson }
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
				fields
				digest
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
			"type": p.dmensObjectType(),
			"fieldJson": map[string]map[string]map[string]interface{}{
				"value": {
					"fields": {
						"action": ACTION_POST,
						"poster": user,
					},
				},
			},
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
