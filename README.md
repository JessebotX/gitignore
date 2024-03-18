# gitignore
Fetch GitHub's `.gitignore` files.

## Install
### Dependencies
- Go <https://go.dev/>

### Installation
```shell
go install github.com/jessebotx/gitignore@latest
```

### Usage
List out available gitignore types that can be fetched.

```shell
gitignore --types
```

### Examples
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
gitignore go node
```

## Permissions/License
See [LICENSE](./LICENSE.txt) for more information.

## See also
- Inspired by (and also a clone of) `msfeldstein/gitignore` on npm <https://github.com/msfeldstein/gitignore>

