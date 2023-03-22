package dmens

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/coming-chat/wallet-SDK/core/base"
)

// BatchQueryUserByAddressJson
// @param jsonString address array's json string. e.g. `["0x1","0x2",   "0x3"]`
func (p *Poster) BatchQueryUserByAddressJson(jsonString string) (*UserPage, error) {
	data := []byte(jsonString)
	var addresses []string
	err := json.Unmarshal(data, &addresses)
	if err != nil {
		return nil, errors.New("invalid json string of address array")
	}
	arrayString := strings.Replace(jsonString, ",", " ", -1)
	return p.batchQueryUserByArrayString(arrayString, len(addresses))
}

func (p *Poster) BatchQueryUserByAddressArray(array *base.StringArray) (*UserPage, error) {
	data, err := json.Marshal(array.Values)
	if err != nil {
		return nil, err
	}
	arrayString := strings.Replace(string(data), ",", " ", -1)
	return p.batchQueryUserByArrayString(arrayString, array.Count())
}

// batchQueryUserByArrayString
// @param str address array's string(no ','). e.g. `["0x1" "0x2"   "0x3"]`
func (p *Poster) batchQueryUserByArrayString(str string, count int) (*UserPage, error) {
	// `filter: {address: {in: ["0x1" "0x2" "0x3"]}}`
	filter := `filter: {address: {in: ` + str + `}}`
	return p.queryUserInfos(count, "", filter)
}

// QueryUserInfoByAddress
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

	user := page.FirstObject()
	if user == nil {
		return nil, nil
	}
	name, _ := p.QuerySuiNameByAddress(user.Address)
	if name != nil {
		user.SuiName = name.Value
	}
	return user, nil
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
			  background
			  website
			  identification
			  item
			}
		  }
		  pageInfo {
			endCursor
			hasNextPage
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

	page := out.MapToUserPage(pageSize)
	p.BatchQueryNFTAvatarForUserPage(page)
	return page, nil
}
