package dmens

import "encoding/json"

// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserFollowers(user string) (string, error) {
	if user == "" {
		user = p.Address
	}
	query := Query{
		Query: `
		query UserFollowers(
			$dmensMetaObjectType: String
			$objectOwner: String
			$profileId: String
		  ) {
			follower(
			  dmensMetaObjectType: $dmensMetaObjectType
			  objectOwner: $objectOwner
			  profileId: $profileId
			) {
			  totalCount
			  nodes {
				followerAddress
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
		Follower json.RawMessage `json:"follower"`
	}
	err := p.makeQueryOut(&query, &out)
	if err != nil {
		return "", err
	}

	return string(out.Follower), nil
}
