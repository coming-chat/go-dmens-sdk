package dmens

func (p *Poster) QueryTrendNoteList(pageSize, offset int) (string, error) {
	query := Query{
		Query: `
		query trendNoteList($type: String, $first: Int, $offset: Int) {
			trendingNotes(
			  filter: { status: { equalTo: "Exists" } }
			  objectType: $type
			  first: $first
			  offset: $offset
			) {
			  totalCount
			  edges {
				cursor
				node {
				  objectId
				  fields
				  createTime
				}
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"type":   p.dmensObjectType(),
			"first":  pageSize,
			"offset": offset,
		},
	}

	var out rawNotePage
	err := p.makeQueryOut(&query, "trendingNotes", &out)
	if err != nil {
		return "", err
	}
	return out.MapToNotePage().JsonString()
}
