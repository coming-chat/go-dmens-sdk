package dmens

// deprated
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) queryUserCharacter(user string) (string, error) {
	if user == "" {
		user = p.Address
	}
	query := Query{
		Query: `
		query UserCharacterByOwner($owner: JSON, $name: JSON) {
			allSuiObjects(
			  filter: {
				status: { equalTo: "Exists" }
				owner: { equalTo: $owner }
				fields: { contains: $name }
			  }
			  orderBy: CREATE_TIME_DESC
			  first: 10
			) {
			  totalCount
			  edges {
				cursor
				node {
				  fields
				}
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

	var out rawUserFollowPage
	err := p.makeQueryOut(&query, "allSuiObjects", &out)
	if err != nil {
		return "", err
	}

	return out.MapToUserPage().JsonString()
}
