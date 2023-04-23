package dmens

import "fmt"

func (p *Poster) QueryTrendNoteList(pageSize int, afterCursor string) (*NotePage, error) {
	filterAppid := ""
	if p.Reviewing {
		filterAppid = fmt.Sprintf("fields: { contains: {value: {fields: {app_id: %v}}}}", appIdForComingChatApp)
	}
	queryString := fmt.Sprintf(`
	query trendNoteList($type: String, $first: Int) {
		trendingNotes(
		  filter: { 
			status: { equalTo: "Exists" } %v
		  }
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
	`, filterAppid)
	query := Query{
		Query: queryString,
		Variables: map[string]interface{}{
			"type":  p.dmensObjectType(),
			"first": pageSize,
		},
		Cursor: afterCursor,
	}

	var out rawNotePage
	err := p.makeQueryOut(&query, "trendingNotes", &out)
	if err != nil {
		return nil, err
	}
	return out.MapToNotePage(p, pageSize), nil
}
