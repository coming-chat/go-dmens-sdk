package dmens

import "github.com/coming-chat/wallet-SDK/core/base"

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

func (p *Poster) IsMyFollowing(address string) (*base.OptionalBool, error) {
	var out []struct {
		Fields struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
	}
	query := p.queryIsFollowing(address)
	err := p.makeQueryOut(query, "allSuiObjects.nodes", &out)
	if err != nil {
		return nil, err
	}

	falseb := &base.OptionalBool{Value: false}
	if len(out) != 1 {
		return falseb, nil
	}
	if out[0].Fields.Name != address {
		return falseb, nil
	}

	return &base.OptionalBool{Value: true}, nil
}
