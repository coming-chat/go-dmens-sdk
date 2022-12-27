package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GraphQLError struct {
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations,omitempty"`
	Message string `json:"message"`
}

func (e GraphQLError) Error() string {
	if len(e.Locations) > 0 {
		return fmt.Sprintf("line %v: %v", e.Locations[0].Line, e.Message)
	} else {
		return e.Message
	}
}

type GraphQLResponse struct {
	Errors []GraphQLError  `json:"errors,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}

// FetchGraphQL [GraphQL](https://graphql.coming.chat/sui-devnet/graphiql)
func FetchGraphQL(query, operationName string, variables map[string]interface{}, graphUrl string, out interface{}) error {
	params := map[string]interface{}{}
	params["query"] = query
	if operationName != "" {
		params["operationName"] = operationName
	}
	if variables != nil {
		params["variables"] = variables
	}
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", graphUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header["Content-Type"] = []string{"application/json"}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resObject := GraphQLResponse{}
	err = json.Unmarshal(respBody, &resObject)
	if err != nil {
		return err
	}
	if resObject.Errors != nil && len(resObject.Errors) > 0 {
		return resObject.Errors[0]
	}
	return json.Unmarshal(resObject.Data, out)
}

// If query has only one statement, `operationName` can be left unspecified
func FetchGraphQLSample(query string, graphUrl string, out interface{}) error {
	return FetchGraphQL(query, "", nil, graphUrl, out)
}
