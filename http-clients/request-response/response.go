/*
Sometimes, we never know the content of a json reponse

We can use map[string]interface{} to handle.

This works since map[string]interface{} is a map of string keys and the values can be any go type
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
any is an alias for interface{}, so UnknownResponseType is also

	map[string]any
*/
type UnknownResponseType = map[string]interface{}

const unknownJsonResponse = `
{
	"users": [
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
	],
	"issues": [
		{
			"id": "1",
			"title": "Issue 1",
			"status": "open",
			"estimate": 1
		}
	]
}
`

func UnknownHandling() {
	var transformedVal UnknownResponseType

	if err := json.Unmarshal([]byte(unknownJsonResponse), &transformedVal); err != nil {
		log.Fatalf("Error unmarshalling unknown JSON: %v", err)
	}

	fmt.Printf("Transformed Value:\n%+v\n", transformedVal)

	fmt.Println("--------------------------------")

	fmt.Printf("Users:\n%+v\n", transformedVal["users"])

	fmt.Printf("First User:\n%+v\n", transformedVal["users"].([]interface{})[0])
}
