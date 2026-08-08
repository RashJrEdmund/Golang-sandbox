## Unit Test

- If you want to add unit tests [see](https://go.dev/doc/tutorial/add-a-test), for using the native `testing` library

Go's built-in support for unit testing makes it easier to test as you go.
Specifically, using naming conventions, Go's testing package, and the go test command,
you can quickly write and execute tests.

- Run the go test command to run tests:
  _<https://pkg.go.dev/cmd/go#hdr-Test_packages>_
  Run this in the directory u want the tests executed. In our case, cd into `internal/auth/` and run

  ```bash
    go test
  ```

  You should see an output like

  ```bash
    st
    PASS
    ok      github.com/RashJrEdmund/go-sandbox/chirpy/internal/auth 0.005s
  ```
