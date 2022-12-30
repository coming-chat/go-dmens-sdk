package dmens

import (
	"fmt"
)

// @param address If the address is empty, the poster's address will be queried.
func (p *Poster) QueryUserInfoByAddress(address string) (string, error) {
	if address == "" {
		address = p.Address
	}
	filter := `filter: {address: {equalTo: "` + address + `"}}`
	return p.queryUserInfos(1, "", filter)
}

func (p *Poster) QueryUsersByName(name string, pageSize int, afterCursor string) (string, error) {
	filter := `filter: {name: {likeInsensitive: "%` + name + `%"}}`
	return p.queryUserInfos(pageSize, afterCursor, filter)
}

func (p *Poster) queryUserInfos(pageSize int, afterCursor string, filter string) (string, error) {
	queryString := fmt.Sprintf(`
	query UserInfos($first: Int) {
		allSuiAddressNames(
		  first: $first
		  after: #cursor#
		  %v
		) {
		  totalCount
		  edges {
			cursor
			node {
			  address
			  avatar
			  bio
			  name
			  nodeId
			}
		  }
		}
	  }
	`, filter)
	query := Query{
		Query: queryString,
		Variables: map[string]interface{}{
			"first": pageSize,
		},
		Cursor: afterCursor,
	}

	var out rawUserPage
	err := p.makeQueryOut(&query, "allSuiAddressNames", &out)
	if err != nil {
		return "", err
	}

	return out.MapToUserPage().JsonString()
}
