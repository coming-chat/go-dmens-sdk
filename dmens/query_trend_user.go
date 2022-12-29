package dmens

import "encoding/json"

func (p *Poster) QueryTrendUserList() (string, error) {
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
			"first":     20,
			"profileId": p.GlobalProfileId,
		},
	}

	var out json.RawMessage
	err := p.makeQueryOut(&query, "trendingCharacters", &out)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
