/*
  JSON is javascript object notation. and go does not have a built in JSON parser.
  so we use the encoding/json package to parse JSON.
  we can use the json.Unmarshal function to parse the JSON into a struct.
  we can use the json.Marshal function to marshal a struct into JSON.
  we can use the json.MarshalIndent function to marshal a struct into JSON with indentation.
  we can use the json.NewDecoder function to create a new JSON decoder.
  we can use the json.NewEncoder function to create a new JSON encoder.
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	// // READING FROM json-decode.go
	// jsonDecode()
	// fmt.Println(Delimiter)

	// // READING FROM json-unmarshal.go
	// jsonUnmarshal()
	// fmt.Println(Delimiter)

	// TestDecodingStringifiedUserJson()
	// fmt.Println(Delimiter)

	// // READING FROM json-marshal.go
	// jsonMarshal()

	fmt.Println(Delimiter)
	UnknownHandling()

	fmt.Println("\nDone")
}

func jsonDecode() {
	decodedIssues, err := DecodeIssues()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Decoded issues:", len(decodedIssues), "\n")
	fmt.Println(decodedIssues)
}

func jsonUnmarshal() {
	unmarshalledIssues, err := UnmarshalIssues()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Unmarshalled issues:", len(unmarshalledIssues), "\n")
	fmt.Println(unmarshalledIssues)
}

func jsonMarshal() {
	TestMarshalUserStruct()
}
