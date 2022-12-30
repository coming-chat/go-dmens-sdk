package dmens

func (p *Poster) QueryTrendUserList(pageSize int) (string, error) {
	query := Query{
		Query: `
		query trendingCharacters($first: Int, $profileId: String) {
			trendingCharacters(first: $first, profileId: $profileId) {
			  nodes {
				followerNumber
				owner
			  }
			  totalCount
			}
		  }
		`,
		Variables: map[string]interface{}{
			"first":     pageSize,
			"profileId": p.GlobalProfileId,
		},
	}

	var out []User
	err := p.makeQueryOut(&query, "trendingCharacters.nodes", &out)
	if err != nil {
		return "", err
	}

	return JsonString(out)
}
