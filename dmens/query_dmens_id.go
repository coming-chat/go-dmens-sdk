package dmens

import (
	"errors"

	"github.com/coming-chat/go-dmens-sdk/graphql"
)

func (p *Poster) QueryDmensObjectId() *Query {
	return &Query{
		Query: `
		query dmensObjectId($owner: JSON, $type: String) {
			allSuiObjects(
			  filter: { owner: { equalTo: $owner }, type: { equalTo: $type } }
			) {
			  nodes {
				objectId
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"owner": map[string]string{
				"AddressOwner": p.Address,
			},
			"type": p.ContractAddress + "::dmens::DmensMeta",
		},
	}
}

func (p *Poster) initialDmensObjecId() error {
	var res = struct {
		AllSuiObjects struct {
			Nodes []struct {
				ObjectId string `json:"objectId"`
			} `json:"nodes"`
		} `json:"allSuiObjects"`
	}{}
	query := p.QueryDmensObjectId()
	err := graphql.FetchGraphQL(query.Query, query.OperationName, query.Variables, p.GraphqlUrl, &res)
	if err != nil {
		return errors.New(err.Error())
	}
	if len(res.AllSuiObjects.Nodes) == 0 {
		return nil
	}
	p.DmensNftId = res.AllSuiObjects.Nodes[0].ObjectId
	return nil
}
