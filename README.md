**You may find some resources in:**

- [resource.md](./docs/resources.md)
- [some postgresql commands](./docs/postgresql-cmds.md)

## Index

1. [My Posts On X](#my-posts-on-x)
2. [Interesting Reads](#interesting-reads)
3. [Others](#others)
4. [JQ](#jq)

## My Posts On X

1 - [Setup and entry points](https://x.com/orashus/status/2059721910754161037) (Nothing much)

2 - [Became a Boot.dev member](https://x.com/orashus/status/2075039939356438694)

3 - [Naked Return Dilemma](https://x.com/orashus/status/2075553448787947675)

4 - [Type assertions](https://x.com/orashus/status/2077069684667543793)

5 - [URL shortener Project](https://x.com/orashus/status/2078129654901198971)

6 - [Interfaces](https://x.com/orashus/status/2079313100776329590)

7 - [Panic](https://x.com/orashus/status/2079589069844054191)

8 - [Variadic and spread operator](https://x.com/orashus/status/2079926280183583104)

9 - [Strings / Unicode](https://x.com/orashus/status/2080426435592737247)

10 - [Pointers](https://x.com/orashus/status/2081988664091930964)

11 - [Goroutines](https://x.com/orashus/status/2083451489037558152)

12 - [Channels](https://x.com/orashus/status/2083577936368566364)

13 - [Escape analysis](https://x.com/orashus/status/2084364744845045981)

## Interesting Reads

- [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

- [The Grug Brained Developer](https://grugbrain.dev/)

## Others

- See the Concepts Readme for more interesting stuff [concepts/README.md](./concepts/README.md)

- Just like the [Zen of Python](https://peps.python.org/pep-0020/),
the [Go Proverbs](https://go-proverbs.github.io/) are a beautiful collection of wise words from Rob Pike, one of Go's creators

```md
  ### Go Proverbs

  ___

  Don't communicate by sharing memory, share memory by communicating.

  Concurrency is not parallelism.

  Channels orchestrate; mutexes serialize.

  The bigger the interface, the weaker the abstraction.

  Make the zero value useful.

  interface{} says nothing.

  Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.

  A little copying is better than a little dependency. // Yup, I know this one from somewhere 😅

  Syscall must always be guarded with build tags.

  Cgo must always be guarded with build tags.

  Cgo is not Go.

  With the unsafe package there are no guarantees.

  Clear is better than clever.

  Reflection is never clear.

  Errors are values.

  Don't just check errors, handle them gracefully.

  Design the architecture, name the components, document the details.

  Documentation is for users.

  Don't panic.
```

## JQ

[jq](https://github.com/jqlang/jq) is a powerful command-line tool for processing JSON data. It's a favorite among developers for working with JSON because it can:

- Parse JSON: easily read and extract data from JSON responses.
- Manipulate JSON: modify JSON data on the fly.
- Filter JSON: find exactly what you're looking for within large JSON structures.

**Example**
Suppose you have a JSON file named user.json with content:

```json
  {
    "name": "John",
    "age": 30,
    "city": "New York"
  }
```

To extract the name field, you would use the [object identifier index](https://jqlang.github.io/jq/manual/#object-identifier-index) like so:

```bash
  jq '.name' user.json
  # "John"
```

To get a field from each element in an array you can use the [array/object value iterator](https://jqlang.github.io/jq/manual/#array-object-value-iterator)
.[], which can in turn be combined with the identifier index like so:

```bash
  curl https://jsonplaceholder.typicode.com/users | jq '.[].username'
  # "Bret"
  # "Antonette"
  # "Samantha"
  # "Karianne"
  # "Kamren"
  # "Leopoldo_Corkery"
  # "Elwyn.Skiles"
  # "Maxime_Nienow"
  # "Delphine"
  # "Moriah.Stanton"
```

Multiple Fields

```bash
  curl https://jsonplaceholder.typicode.com/users/1 | jq '.name, .email'
  # "Leanne Graham"
  # "Sincere@april.biz"
```

- Another typical use is `curl` to parse JSON responses directly from HTTP requests. Eg.

```bash
  curl https://jsonplaceholder.typicode.com/users/1 | jq .username
  # "Bret"
```
