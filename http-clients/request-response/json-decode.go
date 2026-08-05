package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
DECODE JSON:
	We can decode JSON bytes (or strings) into a Go struct using json.Unmarshal or a json.Decoder.

	The Decode method of json.Decoder streams data from an io.Reader into a Go struct,
	while json.Unmarshal works with data that's already in []byte format.

	Using a json.Decoder can be more memory-efficient because it doesn't load all the data into memory at once.
	json.Unmarshal is ideal for small JSON data we already have in memory.

	When dealing with HTTP requests and responses, we will likely use json.Decoder since it works directly with an io.Reader.
*/

// To see how to use json.NewDecoder, follow the GetIssuesForDecoding func
func DecodeIssues() ([]Issue, error) {
	var issues []Issue

	err := GetIssuesForDecoding(&issues) // http request to get issue, passing pointer bcs decoder uses pointer to decode into
	if err != nil {
		return nil, err
	}

	return issues, nil
}

const userList = `
[
	{
		"name": "Rash Dev",
		"role": "Backend Developer",
		"remote": true
	},
	{
		"name": "John Doe",
		"role": "Frontend Developer",
		"remote": false
	}
]
`

type User struct {
	Name   string `json:"name"` // This is the field name in the test userList JSON data
	Role   string `json:"role"`
	Remote bool   `json:"remote"`
}

func TestDecodingStringifiedUserJson() {
	fmt.Println("Decoding stringified user JSON...\n")

	var users []User
	decoder := json.NewDecoder(strings.NewReader(userList)) // if this were a response, we'd use the response body as the input source

	if err := decoder.Decode(&users); err != nil {
		log.Fatalf("Error decoding users: %v", err)
		return
	}

	fmt.Printf("%+v\n", users)
}
