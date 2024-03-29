package dmens

// QueryUserFollowing
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserFollowing(user string, pageSize int, afterCursor string) (*UserPage, error) {
	if user == "" {
		user = p.Address
	}
	query := Query{
		Query: `
		query UserFollowing(
			$dmensMetaObjectType: String
			$objectOwner: String
			$first: Int
		  ) {
			following(
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
	err := p.makeQueryOut(&query, "following", &out)
	if err != nil {
		return nil, err
	}

	page := out.MapToUserPage(pageSize)
	p.BatchQueryNFTAvatarForUserPage(page)
	return page, nil
}

// QueryUserFollowCount
// @param user If the user is empty, the poster's address will be queried.
func (p *Poster) QueryUserFollowCount(user string) (*UserFollowCount, error) {
	if user == "" {
		user = p.Address
	}

	follower, err := p.QueryUserFollowers(user, 0, "")
	if err != nil {
		return nil, err
	}
	following, err := p.QueryUserFollowing(user, 0, "")
	if err != nil {
		return nil, err
	}

	return &UserFollowCount{
		User: user,

		FollowerCount:  follower.TotalCount(),
		FollowingCount: following.TotalCount(),
	}, nil
}
