## GUIDE

Read more about go install docs here <https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies>

## Commands

_Read more about go command here <https://pkg.go.dev/cmd/go>_
_And each go module command here <https://go.dev/ref/mod>_

- ### `go mod init {REMOTE}/{USERNAME}/<package_name>`

_<https://go.dev/ref/mod#go-mod-init>_

To initialize new `go.mod` file in the module

```bash
  go mod init github.com/RashJrEdmund

  # go mod init {REMOTE}/{USERNAME}/hellogo
  # - Where {REMOTE} is your preferred remote source provider (i.e. github.com)
  # - And {USERNAME} is your Git username
```

- ### `go build`

_<https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies>_

It compiles go code into a single, statically linked executable program

Which will be named same as the <package_name> in `go.mod`

To build specific file, run

```bash
  go build <file_name> # creates binary named after the <file_name>
```

### `go run`

_<https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program>_

Simply compiles and runs the binary without emitting the binary

### `go install`

_<https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies>_

The go install command compiles and installs a package or packages on your local machine for your personal usage. It installs the package's compiled binary in the GOBIN directory (cd to `~/go/bin/`). We installed the bootdev cli too there

## A Note on "replace"

_Be aware that using the "replace" keyword like we did in the last assignment is not advised,_
_but can be useful to get up and running quickly. The proper way to create and depend on modules_
_is to publish them to a remote repository. When you do that, the "replace" keyword can be dropped from the go.mod:_
_This only works for local-only development:_

```bash
  replace github.com/RashJrEdmund/go-sandbox/bootdev/mystrings => ../mystrings
```

_If we want the import to work for everyone, we need to make sure the dependency (mystrings in this case) actually exists on  github as a repo. that'll be <https://github.com/RashJrEdmund/mystrings>._
