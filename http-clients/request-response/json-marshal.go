/*
Marshal JSON: https://pkg.go.dev/encoding/json#Unmarshal
	Just as unmarshal converts JSON into a Go struct, marshal converts a Go struct into JSON.

	Marshal is the opposite of unmarshal.

	We still have to write the tags for the fields we want to marshal
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type admin struct {
	User User `json:"user"` // prints {"user":{"name":"Jane Doe","role":"Admin","remote":true}}
}

func TestMarshalUserStruct() {
	admin := admin{
		User: User{
			Name:   "Jane Doe",
			Role:   "Admin",
			Remote: true,
		},
	}

	jsonBytes, err := json.Marshal(admin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(jsonBytes))
}
