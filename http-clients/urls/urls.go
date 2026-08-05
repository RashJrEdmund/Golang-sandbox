/*
DOCS: https://pkg.go.dev/net/url#Parse

The url package provides a way to parse and construct URLs.

It's useful for:
- Parsing URLs
- Constructing URLs
- Querying URLs
- and more!

You can instantiate a URL struct by calling url.Parse with a string URL.
and then your can access the URL's components using the URL struct's fields.
*/

package main

import (
	"log"
	"net/url"

	"fmt"
)

// const urlString = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

/*
There are 8 main parts of a URL, though not all the sections are always present.

	See ParsedURL below for the 8 parts in order.

A COMPLETE URL LOOKS LIKE THIS:

	"protocol://username:password@hostname:port/pathname?search#hash"
*/

/*
INTERESTING:

-> Some urls protocols like http have an authority component that includes the username and password.

	Read here: https://www.rfc-editor.org/info/rfc3986/#section-3.2

-> The port in a URL is a virtual point where network connections are made.

		Ports are managed by a computer's operating system and are numbered from 0 to 65,535
		(Though port 0 is reserved for the system API).

	  - > The port component of a URL is often not visible when browsing normal sites on the internet,
	    because 99% of the time you're using the default ports for the HTTP and HTTPS schemes:
	    80 and 443 respectively.

	    http://google.com:80/?q=orashus.com is equal to http://google.com/?q=orashus.com
*/
const completeUrl = "http://rash_dev:12345@domain.com/path?sort=createdAt#id-section"

type ParsedURL struct {
	protocol string // protocol (http, ftp, mailto, https, etc.) | Required
	username string // username | Optional
	password string // password | Optional
	hostname string // domain name | Required
	port     string // port | Optional | defaults to 80 for http and 443 for https
	pathname string // pathname | Optional | defaults to /
	search   string // query parameters | Optional
	hash     string // fragment | Optional
}

func main() {

	url, err := url.Parse(completeUrl)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
		return
	}

	password, _ := url.User.Password() // returns password and error (if any)

	parsedUrl := ParsedURL{
		protocol: url.Scheme,
		username: url.User.Username(),
		password: password,
		hostname: url.Hostname(),
		port:     url.Port(),
		pathname: url.Path,
		search:   url.Query().Encode(),
		hash:     url.Fragment,
	}

	fmt.Printf("URL: %+v\n", parsedUrl)
}
