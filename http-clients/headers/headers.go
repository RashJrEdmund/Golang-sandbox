/*
Headers are just case-insensitive key-value pairs that pass additional metadata about
	the request or response.

HTTP requests from a web browser automatically carry with them many headers, including but not limited to:
	-> The type of client (e.g. Google Chrome)
	-> The Operating system (e.g. Windows)
	-> The preferred language (e.g. US English)
*/

/*
In Go, the net/http package provides us with the necessary tools to work with HTTP headers.
We can access headers through the Header type, which is essentially a map of string slices (map[string][]string).
*/

package main

import (
	"fmt"
	"net/http"
)

const url = ProjectsUrl

func main() {
	// creating a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error creating request: ", err)
		return
	}

	// setting a header on the new request
	req.Header.Set("x-api-key", "123456789")

	// making the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error making request: ", err)
		return
	}
	defer res.Body.Close()

	// reading a header from the response
	header := res.Header.Get("last-modified")
	fmt.Println("last modified: ", header)

	// deleting a header from the response
	res.Header.Del("last-modified")

	fmt.Println(Delimiter)

	fmt.Println("Content-Type: ", res.Header.Get("Content-Type"))

	fmt.Println(Delimiter)
	fmt.Println("headers: ", res.Header)
}
