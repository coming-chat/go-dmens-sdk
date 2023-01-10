package dmens

func (p *Poster) QueryTrendUserList(pageSize int) (*UserPage, error) {
	query := Query{
		Query: `
		query trendingCharacters($first: Int, $profileId: String) {
			trendingCharacters(first: $first, profileId: $profileId) {
			  totalCount
			  edges {
				cursor
				node {
				  fields
				  followerNumber
				  owner
				}
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"first":     pageSize,
			"profileId": p.GlobalProfileTableId,
		},
	}

	var out rawUserFollowPage
	err := p.makeQueryOut(&query, "trendingCharacters", &out)
	if err != nil {
		return nil, err
	}

	return out.MapToUserPage(pageSize), nil
}
