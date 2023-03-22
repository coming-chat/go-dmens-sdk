package dmens

// QueryUserFollowers
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserFollowers(user string, pageSize int, afterCursor string) (*UserPage, error) {
	if user == "" {
		user = p.Address
	}
	query := Query{
		Query: `
		query UserFollowers(
			$dmensMetaObjectType: String
			$objectOwner: String
			$first: Int
		  ) {
			follower(
			  dmensMetaObjectType: $dmensMetaObjectType
			  objectOwner: $objectOwner
			  first: $first
			  after: #cursor#
			) {
			  totalCount
			  edges {
				cursor
				node {
				  address
			 	  avatar
				  bio
				  name
				  nodeId
				  background
				  website
				  identification
				  item
				}
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"dmensMetaObjectType": p.dmensMetaObjectType(),
			"objectOwner":         user,
			"first":               pageSize,
		},
		Cursor: afterCursor,
	}

	var out rawUserPage
	err := p.makeQueryOut(&query, "follower", &out)
	if err != nil {
		return nil, err
	}

	page := out.MapToUserPage(pageSize)
	p.BatchQueryNFTAvatarForUserPage(page)
	return page, nil
}
