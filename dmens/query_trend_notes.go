package dmens

func (p *Poster) QueryTrendNoteList(pageSize int, afterCursor string) (string, error) {
	query := Query{
		Query: `
		query trendNoteList($type: String, $first: Int) {
			trendingNotes(
			  filter: { status: { equalTo: "Exists" } }
			  objectType: $type
			  first: $first
			  after: #cursor#
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
			"type":  p.dmensObjectType(),
			"first": pageSize,
		},
		Cursor: afterCursor,
	}

	var out rawNotePage
	err := p.makeQueryOut(&query, "trendingNotes", &out)
	if err != nil {
		return "", err
	}
	return out.MapToNotePage().JsonString()
}
