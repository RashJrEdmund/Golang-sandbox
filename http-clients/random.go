// USING THIS TO RUN RANDOM HTTP CLIENT CODE AND SEE THE OUTPUT

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {
	issueBytes, err := getIssueData(issueURL)
	if err != nil {
		log.Fatalf("Error getting issue data: %v", err)
	}
	prettyData, err := prettify(issueBytes)
	if err != nil {
		log.Fatalf("Error prettifying data: %v", err)
	}
	fmt.Println(prettyData)
}

/*
This is from the lessons:

	-> https://www.boot.dev/lessons/07c71a3f-742b-40c6-a775-f289eb417190
	-> https://www.boot.dev/lessons/0deb4af3-1f09-4676-bba9-912b8347632e
*/
func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}

	return data, nil
}

func prettify(byteData []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, byteData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}
