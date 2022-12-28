package dmens

import (
	"encoding/json"
	"errors"

	"github.com/coming-chat/go-dmens-sdk/graphql"
)

type Query struct {
	Query         string
	OperationName string
	Variables     map[string]interface{}
}

func (p *Poster) MakeQuery(q *Query) (string, error) {
	var out json.RawMessage
	err := p.makeQueryOut(q, &out)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (p *Poster) makeQueryOut(q *Query, out interface{}) error {
	err := graphql.FetchGraphQL(q.Query, q.OperationName, q.Variables, p.GraphqlUrl, out)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
