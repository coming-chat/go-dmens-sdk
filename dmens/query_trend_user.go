package dmens

func (p *Poster) QueryTrendUserList(pageSize int) (*UserPage, error) {
	query := Query{
		Query: `
		query trendingCharacters($first: Int) {
			trendingCharacters(first: $first) {
			  totalCount
			  edges {
				cursor
				node {
				  address
				  avatar
				  bio
				  name
				  followerNumber
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
			"first": pageSize,
		},
	}

	var out rawUserFollowPage
	err := p.makeQueryOut(&query, "trendingCharacters", &out)
	if err != nil {
		return nil, err
	}

	page := out.MapToUserPage(pageSize)
	p.BatchQueryNFTAvatarForUserPage(page)
	return page, nil
}
