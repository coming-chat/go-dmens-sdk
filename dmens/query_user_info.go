package dmens

import (
	"encoding/json"
	"fmt"
)

// @param address If the address is empty, the poster's address will be queried.
func (p *Poster) QueryUserInfoByAddress(address string) (string, error) {
	if address == "" {
		address = p.Address
	}
	filter := `filter: {address: {equalTo: "` + address + `"}}`
	return p.queryUserInfos(1, 0, filter)
}

func (p *Poster) QueryUsersByName(name string, pageSize, offset int) (string, error) {
	filter := `filter: {name: {likeInsensitive: "%` + name + `%"}}`
	return p.queryUserInfos(pageSize, offset, filter)
}

func (p *Poster) queryUserInfos(pageSize, offset int, filter string) (string, error) {
	queryString := fmt.Sprintf(`
	query UserInfos($first: Int, $offset:Int) {
		allSuiAddressNames(
		  first: $first
		  offset: $offset
		  %v
		) {
		  nodes {
			address
			avatar
			bio
			name
			nodeId
		  }
		  pageInfo {
			hasNextPage
		  }
		  totalCount
		}
	  }
	`, filter)
	query := Query{
		Query: queryString,
		Variables: map[string]interface{}{
			"first":  pageSize,
			"offset": offset,
		},
	}

	var out json.RawMessage
	err := p.makeQueryOut(&query, "allSuiAddressNames", &out)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
