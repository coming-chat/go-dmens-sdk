package dmens

import "encoding/json"

func (p *Poster) QueryTrendTwittersList(pageSize, offset int) (string, error) {
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
			"first":  pageSize,
			"offset": offset,
		},
	}

	var out json.RawMessage
	err := p.makeQueryOut(&query, "trendingNotes", &out)
	if err != nil {
		return "", err
	}

	res := string(out)
	return res, nil
}
