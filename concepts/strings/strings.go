/*
### Understanding Unicode, UTF-8 and string encoding
Joel Spolsky's article on the "Absolute Minimum Every Software Developer Absolutely Positively Must Know About Unicode And Character Sets (no excuses!)"
https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/

X POST
https://x.com/orashus/status/2080426435592737247

### CODE POINTS, CHARACTERS, AND RUNES
DOCS: https://go.dev/blog/strings
Strings are immutable sequences of bytes.
	- The Go language strings are implemented as immutable byte slices.
	- The bytes are stored in a contiguous block of memory, and the string is a pointer to the first byte.
	- The string is immutable, so any changes to the string will create a new string.
	- The string is a value type, so it is passed by value.
	- Strings are built from bytes so indexing them yields bytes, not characters

Go's standard library provides strong support for interpreting UTF-8 text.
	- The most important such package is unicode/utf8, which contains helper routines to validate,
		disassemble, and reassemble UTF-8 strings.
	- DOCS: https://pkg.go.dev/unicode/utf8

Read this to get some understanding about RUNES: https://go.dev/blog/strings#code-points-characters-and-runes
	- A Rune is a Unicode code point.
	- A Unicode code point is a 32-bit value that represents a Unicode code point.
	- The Go language defines the word rune as an alias for the type int32, so programs can be clear when
	 	an integer value represents a code point.

When you need to decode a string into its individual code points, you can use the utf8.DecodeRuneInString function.
	- DOCS: https://pkg.go.dev/unicode/utf8#DecodeRuneInString

When you need to encode a string into its individual code points, you can use the utf8.EncodeRune function.
	- DOCS: https://pkg.go.dev/unicode/utf8#EncodeRune
*/

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// %#U - Shows the code point’s Unicode value and its printed representation like this: U+0041 'A'

func main() {
	for _, r := range "Hello, 世界" {
		v, size := utf8.DecodeRuneInString(string(r))
		fmt.Printf("%q --- %d --- %#U --- %d --- %q\n", r, r, r, size, v)
	}

	fmt.Printf("Bytes: %v\n", len("Hi 🦫"))

	const word = "HiA012"

	fmt.Printf("Word Contains \"Hi\" %t\n", strings.Contains(strings.ToLower(word), "hi"))

	for i, rune := range []rune(word) {
		fmt.Printf("%q %v %T %v\n", rune, rune, rune, word[i])
	}

	fmt.Printf("%v %v\n", []rune(word), int('A'))

	// runForRangeLoop()
}

// func runForRangeLoop() {
// 	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

// 	fmt.Println("sample:\n\t", sample)
// 	const sample2 = "azAZHello world!"

// 	// 65 - 90: A-Z
// 	// 97 - 122: a-z
// 	// 48 - 57: 0-9
// 	// 32 - 126: printable characters
// 	// 32 - 126: printable characters

// 	for _, s := range sample2 {
// 		fmt.Printf("sample2:\n\t %%v: %v -> %%q: %q -> %%c: %c\n", s, s, s)
// 	}
// }
