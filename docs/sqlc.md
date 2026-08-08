
## SQLC

These and the goose.md docs are written to reference the chirpy projects

_<https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html>_
_<https://sqlc.dev>_

**Checkout the storage course on http-servers on boot.dev to find more if not satisfied**

SQLC is a Go program that generates Go code from SQL queries.
It's not exactly an ORM, but rather a tool that makes working with raw SQL easy and type-safe.

- We can be using Goose to manage our database migrations (the schema), and then using SQLC to generate Go code that our application can use to interact with the database (run queries).

### Installing using go

_<https://docs.sqlc.dev/en/latest/overview/install.html>_

  ```bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  ```

### Configuring using a yaml file

_<https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html>_

  ```yml
    version: "2"
    sql:
      - schema: "sql/schema"
        queries: "sql/queries"
        engine: "postgresql"
        gen:
          go:
            out: "internal/database"
  ```

  We're telling SQLC to look in the sql/schema directory for our schema structure (which is the same set of files that Goose uses, but sqlc automatically ignores "down" migrations), and in the sql/queries directory for queries. We're also telling it to generate Go code in the internal/database directory.

### Generate Type safe code from schemas

Generate the Go code. Run `sqlc generate` from the root of your project.
It should create a new package of go code in `internal/database` directory

```bash
  sqlc generate # to generate type-safe code
```

- You'll notice that the generated code relies on [Google's uuid package](https://pkg.go.dev/github.com/google/uuid),
so you'll need to add that to your module:

  ```bash
    go get github.com/google/uuid
  ```

- We need to add and import a Postgres driver so our program knows how to talk to the database.
Install it in your module as well:

  ```bash
    go get github.com/lib/pq
  ```

  Add this import to the top of your main.go file.
  You have to import the driver, but you don't use it directly anywhere in your code. The underscore tells Go that you're importing it for its side effects, not because you need to use it:

  ```go
    import _ "github.com/lib/pq"
  ```
