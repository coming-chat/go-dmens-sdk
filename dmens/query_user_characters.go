package dmens

import "encoding/json"

// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserCharacter(user string) (string, error) {
	if user == "" {
		user = p.Address
	}
	query := Query{
		Query: `
		query UserCharacterByOwner($first: Int = 10, $owner: JSON, $name: JSON) {
			allSuiObjects(
			  filter: {
				status: { equalTo: "Exists" }
				owner: { equalTo: $owner }
				fields: { contains: $name }
			  }
			  orderBy: CREATE_TIME_DESC
			  first: $first
			) {
			  totalCount
			  nodes {
				createTime
				dataType
				digest
				fields
				isUpdate
				nodeId
				hasPublicTransfer
				owner
				objectId
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
			"owner": map[string]string{
				"ObjectOwner": p.GlobalProfileId,
			},
			"name": map[string]string{
				"name": user,
			},
		},
	}

	var out json.RawMessage
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
