package dmens

import "encoding/json"

// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserFollowing(user string) (string, error) {
	if user == "" {
		user = p.Address
	}
	query := Query{
		Query: `
		query UserFollowing(
			$dmensMetaObjectType: String
			$objectOwner: String
			$profileId: String
		  ) {
			following(
			  dmensMetaObjectType: $dmensMetaObjectType
			  objectOwner: $objectOwner
			  profileId: $profileId
			) {
			  totalCount
			  nodes {
				followingAddress
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"dmensMetaObjectType": p.dmensMetaObjectType(),
			"objectOwner":         user,
			"profileId":           p.GlobalProfileId,
		},
	}

	var out struct {
		Following json.RawMessage `json:"following"`
	}
	err := p.makeQueryOut(&query, &out)
	if err != nil {
		return "", err
	}

	return string(out.Following), nil
}
