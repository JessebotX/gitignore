# gitignore
Fetch GitHub's `.gitignore` files.

## Installation: CLI Tool
### Dependencies
- Go <https://go.dev/>

### Steps
```shell
go install github.com/JessebotX/gitignore/cmd/gitignore@latest
```

### Usage
List out available gitignore types that can be fetched.

```shell
gitignore --types
```

### CLI Examples
Get list of all gitignore types that are supported

```shell
gitignore --types
```

Create a .gitignore file for a Go project

```shell
gitignore go
```

Create a .gitignore file for a Python project

```shell
gitignore python
```

Create a .gitignore file for a Go and a Node.js project

```shell
gitignore go node
```

Print out .gitignore for a Go and a Node.js project into stdout

```shell
gitignore go node --print
```

## Permissions/License
See [LICENSE](./LICENSE.txt) for more information.

## See also
- Inspired by (and also a clone of) `msfeldstein/gitignore` on npm <https://github.com/msfeldstein/gitignore>

