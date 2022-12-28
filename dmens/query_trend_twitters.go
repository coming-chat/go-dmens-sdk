package dmens

import "encoding/json"

func (p *Poster) QueryTrendTwittersList(first, offset int) (string, error) {
	query := Query{
		Query: `
		query trendTwittersList($type: String, $first: Int, $offset: Int) {
			trendingNotes(
			  filter: { status: { equalTo: "Exists" } }
			  objectType: $type
			  first: $first
			  offset: $offset
			) {
			  nodes {
				objectId
				owner
				fields
				dataType
				createTime
				digest
				hasPublicTransfer
				isUpdate
				previousTransaction
				status
				storageRebate
				type
				updateTime
				version
				actionNumber
			  }
			  totalCount
			}
		  }
		`,
		Variables: map[string]interface{}{
			"type":   p.dmensObjectType(),
			"first":  first,
			"offset": offset,
		},
	}

	var out struct {
		TrendingNotes json.RawMessage `json:"trendingNotes"`
	}
	err := p.makeQueryOut(&query, &out)
	if err != nil {
		return "", err
	}

	res := string(out.TrendingNotes)
	return res, nil
}
