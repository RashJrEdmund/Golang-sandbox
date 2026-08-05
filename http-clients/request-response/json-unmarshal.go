/*
Unmarshal JSON:
	We can decode JSON bytes (or strings) into a Go struct using json.Unmarshal or a json.Decoder.

	The Decode method of json.Decoder streams data from an io.Reader into a Go struct,
	while json.Unmarshal works with data that's already in []byte format.

	Using a json.Decoder can be more memory-efficient because it doesn't load all the data into memory at once.
	json.Unmarshal is ideal for small JSON data we already have in memory.

	When dealing with HTTP requests and responses, we will likely use json.Decoder since it works directly with an io.Reader.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func UnmarshalIssues() ([]Issue, error) {
	arrIssueBytes, err := GetIssuesForUnmarshalling() // http request to get issue
	if err != nil {
		return nil, err
	}

	var issues []Issue
	if err := json.Unmarshal(arrIssueBytes, &issues); err != nil {
		return nil, fmt.Errorf("Error unmarshalling issues: %w", err)
	}

	return issues, nil
}
