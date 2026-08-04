/*
TypeScript does have a pretty incredible type system. Here's one of the things it can do:

	type SendingChannel = "email" | "sms" | "phone";

	function sendNotification(ch: SendingChannel, message: string) {
		// send the message
	}

This sendingChannel type that we've created is a union type. It can only be one of the three strings that we've defined.
That means when a developer calls sendNotification() they can't accidentally pass an invalid sendingChannel like "slack"
or even a misspelled "emil". The TypeScript compiler will catch that mistake at compile time.

The closest thing we have to a union type in Go is a type definition: https://go.dev/ref/spec#Type_definitions

	type SendingChannel string

	const (
		Email SendingChannel = "email"
		SMS   SendingChannel = "sms"
		Phone SendingChannel = "phone"
	)

	func sendNotification(ch SendingChannel, message string) {
		// send the message
	}

It's a bit safer than using a plain string in Go, but it's not completely safe. Go will stop us from doing this:

	sendingCh := "slack"
	sendNotification(sendingCh, "hello") // string is not sendingChannel

But it will not stop us from doing this:

	// "slack" is automatically implied as a sendingChannel
	sendNotification("slack", "hello") // string is not sendingChannel

And will also not stop us from converting like this:

	sendingCh := "slack"
	convertedSendingCh := SendingChannel(sendingCh)
	sendNotification(convertedSendingCh, "hello")

The SendingChannel type is just a wrapper for string, and because we made some constants of that type,
most developers will just use those constants
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	checkPermission(Read)

	checkPermission(Perm(Admin)) // this should not be allowed.
}

type Perm string

const (
	Read  Perm = "read"
	Write Perm = "write"
	Exec  Perm = "execute"
)

var Admin = "admin"
var User = Perm("user")

func checkPermission(p Perm) {
	// check the Permission
}
