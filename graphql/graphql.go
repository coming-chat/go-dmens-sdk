package graphql

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Error struct {
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations,omitempty"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	if len(e.Locations) > 0 {
		return fmt.Sprintf("line %v: %v", e.Locations[0].Line, e.Message)
	} else {
		return e.Message
	}
}

type Response struct {
	Errors []Error         `json:"errors,omitempty"`
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
	resObject := Response{}
	err = json.Unmarshal(respBody, &resObject)
	if err != nil {
		return err
	}
	if resObject.Errors != nil && len(resObject.Errors) > 0 {
		basicError := errors.New(resObject.Errors[0].Error())
		// If not converted to a basic error, it may cause Android to crash
		return basicError
	}
	return json.Unmarshal(resObject.Data, out)
}

// FetchGraphQLSample
// If the query has only one statement, `operationName` can be left unspecified
func FetchGraphQLSample(query string, graphUrl string, out interface{}) error {
	return FetchGraphQL(query, "", nil, graphUrl, out)
}
