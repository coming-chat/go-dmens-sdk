package dmens

import (
	"fmt"
)

// @param address If the address is empty, the poster's address will be queried.
func (p *Poster) QueryUserInfoByAddress(address string) (*UserInfo, error) {
	if address == "" {
		address = p.Address
	}
	filter := `filter: {address: {equalTo: "` + address + `"}}`
	page, err := p.queryUserInfos(1, "", filter)
	if err != nil {
		return nil, err
	}
	return page.FirstObject(), nil
}

func (p *Poster) QueryUsersByName(name string, pageSize int, afterCursor string) (*UserPage, error) {
	filter := `filter: {name: {likeInsensitive: "%` + name + `%"}}`
	return p.queryUserInfos(pageSize, afterCursor, filter)
}

func (p *Poster) queryUserInfos(pageSize int, afterCursor string, filter string) (*UserPage, error) {
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
		return nil, err
	}

	return out.MapToUserPage(), nil
}
