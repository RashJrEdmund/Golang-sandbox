/*
In go we make requests using the "net/http" package.
The http.Get function is a convenience function that creates a new GET request and returns the response.
The http.Get function uses the http.DefaultClient under the hood to make the request.
*/

/*
-> http.Get uses the http.DefaultClient to make a request to the given url
-> res is the HTTP response that comes back from the server
-> defer res.Body.Close() ensures that the response body is properly closed after reading. Not doing so can cause memory issues.
-> io.ReadAll reads the response body into a slice of bytes []byte called data
*/

/*
THERE ARE TWO BASIC WAYS TO MAKE A REQUEST IN GO.

	-> The simple but less powerful way:
		http.Get

	-> The verbose but more powerful way:
		http.Client, http.NewRequest, and http.Client.Do
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const url = UsersUrl

func main() {
	fmt.Println(Delimiter)
	loadUsers()

	fmt.Println(Delimiter)

	newComment := Comment{
		Id:      "1",
		UserId:  "1",
		Comment: "This is a comment",
	}

	createComment(
		url,
		"my-api-key-value-goes-here",
		newComment,
	)
}

func loadUsers() {
	// creating a new request

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("x-api-key", "my-api-key-value-goes-here")

	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var users []interface{}
	if err := decoder.Decode(&users); err != nil {
		log.Fatalf("Error decoding response: %v", err)
		return
	}

	fmt.Printf("%+v\n", users)
}

type Comment struct {
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func createComment(url, apiKey string, commentStruct Comment) {
	// encode our comment as json
	jsonData, err := json.Marshal(commentStruct)
	if err != nil {
		log.Fatalf("Error marshalling comment: %v", err)
		return
	}

	// create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData)) // jsonData is the body of the request
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// create a new client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return
	}
	defer res.Body.Close()

	// decode the json data from the response
	// into a new Comment struct
	var comment Comment
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&comment)
	if err != nil {
		log.Fatalf("Error decoding response: %v", err)
		return
	}

	fmt.Printf("%+v\n", comment)
}
