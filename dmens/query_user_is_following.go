package dmens

func (p *Poster) queryIsFollowing(address string) *Query {
	return &Query{
		Query: `
query isMyfollowing($owner: JSON, $fields: JSON) {
  allSuiObjects(
    filter: {owner: {equalTo: $owner}, fields: {contains: $fields}, status: {equalTo: "Exists"}}
  ) {
    nodes {
	  fields
    }
  }
}
		`,
		Variables: map[string]interface{}{
			"owner": p.followsObjectOwner(),
			"fields": map[string]string{
				"name":  address,
				"value": address,
			},
		},
	}
}

func (p *Poster) IsMyFollowing(address string) (bool, error) {
	var out []struct {
		Fields struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
	}
	query := p.queryIsFollowing(address)
	err := p.makeQueryOut(query, "allSuiObjects.nodes", &out)
	if err != nil {
		return false, err
	}

	if len(out) != 1 {
		return false, nil
	}

	if out[0].Fields.Name != address {
		return false, nil
	}
	return true, nil
}
