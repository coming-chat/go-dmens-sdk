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

	return out.MapToUserPage(pageSize), nil
}
