package dmens

import "github.com/coming-chat/wallet-SDK/core/base"

func (p *Poster) IsMyFollowing(address string) (*base.OptionalBool, error) {
	status, err := p.batchQueryIsFollowingStatus("", []string{address})
	if err != nil {
		return nil, err
	}
	b := status[address]
	return &base.OptionalBool{Value: b}, nil
}

// BatchQueryIsFollowingStatus
// Batch query the following status of all users in a specified list.
// The query results will be modified directly in the pointer object.
func (p *Poster) BatchQueryIsFollowingStatus(users *UserPage) error {
	userList := make([]string, users.TotalCount())
	for idx, user := range users.Items {
		userList[idx] = user.Address
	}
	status, err := p.batchQueryIsFollowingStatus("", userList)
	if err != nil {
		return err
	}
	for _, user := range users.Items {
		user.IsFollowing = status[user.Address]
	}
	return nil
}

func (p *Poster) batchQueryIsFollowingStatus(viewer string, users []string) (map[string]bool, error) {
	if viewer == "" {
		viewer = p.Address
	}
	query := Query{
		Query: `
		query BatchUserFollowingStatus(
			$dmensMetaObjectType: String
			$objectOwner: String
			$users: [String!]
		  ) {
			following(
			  dmensMetaObjectType: $dmensMetaObjectType
			  objectOwner: $objectOwner
			  filter: {address: {in: $users}}
			) {
			  edges {
				node {
				  address
				}
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"dmensMetaObjectType": p.dmensMetaObjectType(),
			"objectOwner":         viewer,
			"users":               users,
		},
	}

	var out rawUserPage
	err := p.makeQueryOut(&query, "following", &out)
	if err != nil {
		return nil, err
	}

	status := make(map[string]bool)
	for _, item := range out.Edges {
		status[item.Node.Address] = true
	}
	return status, nil
}
