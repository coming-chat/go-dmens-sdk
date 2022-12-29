package dmens

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coming-chat/go-dmens-sdk/graphql"
)

type Query struct {
	Query         string
	OperationName string
	Variables     map[string]interface{}
}

func (p *Poster) MakeQuery(q *Query) (string, error) {
	var out json.RawMessage
	err := p.makeQueryOut(q, "", &out)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (p *Poster) makeQueryOut(q *Query, path string, out interface{}) error {
	if path == "" {
		return graphql.FetchGraphQL(q.Query, q.OperationName, q.Variables, p.GraphqlUrl, out)
	}

	var oo interface{}
	err := graphql.FetchGraphQL(q.Query, q.OperationName, q.Variables, p.GraphqlUrl, &oo)
	if err != nil {
		return err
	}

	var (
		m  map[string]interface{}
		ok bool
	)
	paths := strings.Split(path, ".")
	for _, path := range paths {
		if path == "" {
			continue
		}
		if m, ok = oo.(map[string]interface{}); ok {
			if oo, ok = m[path]; ok {
				continue
			}
		}
		return fmt.Errorf("Response parsing error, path '%v' not found.", path)
	}

	data, err := json.Marshal(oo)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}

func (p *Poster) dmensMetaObjectType() string {
	return p.ContractAddress + "::dmens::DmensMeta"
}

func (p *Poster) dmensObjectType() string {
	return "0x2::dynamic_field::Field<u64, " + p.ContractAddress + "::dmens::Dmens>"
}

func (p *Poster) addressOwner() map[string]string {
	return map[string]string{
		"AddressOwner": p.Address,
	}
}
