/*
When we receive JSON data in the body of an HTTP response, it comes as a stream of bytes
To decode this JSON into it's appropriate form, we need to know the JSON fields and their types.
The standard encoding/json package uses tags to map JSON fields to struct fields.
*/

/*
	-> defer res.Body.Close() ensures that the response body is properly closed after reading.
		-> Not doing so can cause memory issues.
		-> This is particularly important for large responses.

	-> On the backend, the server is responsible for closing the request body.
		-> so no need for defer req.Body.Close() over there

	-> io.ReadAll reads the response body into a slice of bytes []byte called data
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
When using a decoder or unmarshaler,
	-> the JSON data is decoded into the struct, based on the struct tags.

		`json:"id"` is the tag for the id field // see Issue struct below.

	-> We pass the decoder or unmarshaler the pointer to the struct we want to decode into.

***NOTE***:
	-> The struct fields must be exported (capitalized) to decode JSON.
*/

type Issue struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Estimate int    `json:"estimate"`
}

const url = IssuesUrl

func GetIssuesForUnmarshalling() ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}

	return data, nil
}

/*
IT SHOULD BE NOTED THAT WHEN USING A DECODER AFTER USING http.Get,
THE DECODER CAN READ THE RESPONSE BODY DIRECTLY,
INSTEAD OF USING io.ReadAll AND HAVING TO CONVERT AGAIN WITH

	decoder := json.NewDecoder(strings.NewReader(string(arrIssueBytes)))
*/

func GetIssuesForDecoding(issues *[]Issue) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error making request: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	// var issues []Issue
	if err := decoder.Decode(issues); err != nil {
		return fmt.Errorf("Error decoding issues: %w", err)
	}

	return nil
}
