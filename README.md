# pairedbrackets <br> [![go.mod version][go-img]][go-url] [![CI][ci-img]][ci-url] [![Codecov][codecov-img]][codecov-url] [![Codebeat][codebeat-img]][codebeat-url] [![Maintainability][codeclimate-img]][codeclimate-url] [![Go Report Card][goreportcard-img]][goreportcard-url] [![License][license-img]][license-url] [![Go Reference][godoc-img]][godoc-url]

Linter checks formatting of paired brackets (inspired by [this article](https://www.yegor256.com/2014/10/23/paired-brackets-notation.html)).


## Rule

According to the original [notation](https://www.yegor256.com/2014/10/23/paired-brackets-notation.html), "a bracket should either start/end a line or be paired on the same line".  
With modification for multiline items, the following cases are allowed:
1. Both brackets and all items are in one line.
   ```go
   fmt.Printf("%s, %s!", "Hello", "world")
   ```
1. Left (opening) bracket is the last character on a line and right (closing) bracket starts a new line.
   ```go
   fmt.Printf( // comments and whitespaces are ignored
   	"%s, %s!", "Hello", "world",
   )
   ```
1. If the last item is multiline, it can start on the same line with the left bracket.  
   In this case, the right bracket should be on the same line where the last item ends.     
   ```go
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   	...
   })
   ```


## Linter reports (wordings):

> x.go:1:16: left parenthesis should either be the last character on a line or be on the same line with the last argument
```go
               ⬇
http.HandleFunc("/",
	func(w http.ResponseWriter, r *http.Request) {
		...
	},
)
```

<br>

> x.go:4:1: right parenthesis should be on the previous line
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	...
},
)
⬆
```

<br>

> x.go:5:3: right parenthesis should be on the next line
```go
http.HandleFunc(
	"/",
	func(w http.ResponseWriter, r *http.Request) {
		...
	})
	 ⬆
```


## Examples

<details><summary>Function/method call</summary>
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
fmt.Printf("%s, %s!",
	"Hello", "world")
```
```go
fmt.Printf("%s, %s!",
	"Hello", "world",
)
```
```go
fmt.Printf(
	"%s, %s!",
	"Hello", "world")
```
```go
fmt.Printf("%s %s", "Last", `item
is multiline`,
)
```

</td><td>

```go
fmt.Printf("%s, %s!", "Hello", "world")
```
```go
fmt.Printf(
	"%s, %s!", "Hello", "world",
)
```
```go
fmt.Printf(
	"%s, %s!",
	"Hello", "world",
)
```
```go
fmt.Printf("%s %s", "Last", `item
is multiline`)
```

</td></tr>
</tbody></table>

Functions from `github.com/stretchr/testify/assert` and `github.com/stretchr/testify/require` are ignored by default (see [config](#config)).

</details>

<details><summary>Composite literal</summary>
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
foo := []int{1,
	2, 3}
```
```go
foo := []int{1,
	2, 3,
}
```
```go
foo := []int{
	1,
	2,
	3}
```
```go
foo := []string{"Last", "item", `is
multiline`,
}
```

</td><td>

```go
bar := []int{1, 2, 3}
```
```go
bar := []int{
	1,
	2,
	3,
}
```
```go
bar := []int{
	1, 2, 3,
}
```
```go
bar := []string{"Last", "item", `is
multiline`}
```

</td></tr>
</tbody></table>
</details>

<details><summary>Function parameters</summary>
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func Foo(a int,
	b string, c bool) {
	...
}
```
```go
func Foo(a int,
	b string, c bool,
) {
	...
}
```
```go
func Foo(
	a int,
	b string,
	c bool) {
	...
}
```
```go
func Foo(a int, b string,
) {
	...
}
```
```go
func Foo(a int, b struct {
	X int
	Y string
},
) {
	...
}
```

</td><td>

```go
func Bar(a int, b string) {
	...
}
```
```go
func Bar(
	a int,
	b string,
	c bool,
) {
	...
}
```
```go
func Bar(
	a int, b string, c bool,
) {
	...
}
```
```go
func Bar(a int, b struct {
	X int
	Y string
}) {
	...
}
```

</td></tr>
</tbody></table>
</details>

<details><summary>Function type parameters (generics)</summary>
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func Foo[T int,
	V string]() {
	...
}
```
```go
func Foo[T int,
	V string,
]() {
	...
}
```
```go
func Foo[
	T int,
	V string]() {
	...
}
```
```go
func Foo[T int, V string,
]() {
	...
}
```
```go
func Foo[T int, V interface {
	int | string
},
]() {
	...
}
```

</td><td>

```go
func Bar[T int, V string]() {
	...
}
```
```go
func Bar[
	T int,
	V string,
]() {
	...
}
```
```go
func Bar[
	T int, V string,
]() {
	...
}
```
```go
func Bar[T int, V interface {
	int | string
}]() {
	...
}
```

</td></tr>
</tbody></table>
</details>

<details><summary>Function returns (output parameters)</summary>
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func Foo() (int,
	error) {
	...
}
```
```go
func Foo() (int,
	error,
) {
	...
}
```
```go
func Foo() (
	int,
	error) {
	...
}
```
```go
func Foo() (int, error,
) {
	...
}
```
```go
func Foo() (int, interface {
	Error()
},
) {
	...
}
```

</td><td>

```go
func Bar() (int, error) {
	...
}
```
```go
func Bar() (
	int,
	error,
) {
	...
}
```
```go
func Bar() (
	int, error,
) {
	...
}
```
```go
func Bar() (int, interface {
	Error()
}) {
	...
}
```

</td></tr>
</tbody></table>
</details>


## Other tools

### gofmt

`gofmt` fixes many cases, which `pairedbrackets` complains about. But not all of them. All examples above formatted correctly according to `gofmt`.

### gofumpt

`gofumpt` is just a slightly better than `gofmt`. It fixes some composite literal examples above. But not all of them, and it doesn't fix other examples.


## Usage


The best way is to use [golangci-lint](https://golangci-lint.run/).  
It includes [pairedbrackets](https://golangci-lint.run/usage/linters/#list-item-pairedbrackets) and a lot of other great linters.

### Install

See [official site](https://golangci-lint.run/usage/install/).

### Enable

`pairedbrackets` is disabled by default.  
To enable it, add the following to your `.golangci.yml`:

```yaml
linters:
  enable:
     pairedbrackets
```

### Config

Here is available configuration:
```yaml
linters-settings:
  pairedbrackets:
    # List of regexp patterns of fully qualified function calls to ignore.
    # Example of fully qualified function: github.com/stretchr/testify/require.Equal
    # Example of fully qualified method: (*github.com/stretchr/testify/assert.Assertions).Equal
    # Default: ["github.com/stretchr/testify/assert", "github.com/stretchr/testify/require"]
    ignore-func-calls:
      - github.com/stretchr/testify/assert
      - github.com/stretchr/testify/require
      - ^fmt\.
```

### Run

```shell
golangci-lint run
```


## Usage as standalone linter

### Install

```shell
go install github.com/maratori/pairedbrackets@latest
```

### Run

```shell
pairedbrackets ./...
```

### Flag `-ignore-func-calls`

Flag `-ignore-func-calls` is a comma separated list of regexp patterns of fully qualified function calls to ignore. Default is `github.com/stretchr/testify/assert,github.com/stretchr/testify/require`.

Fully qualified function examples:
- Function: `github.com/stretchr/testify/require.Equal`
- Method: `(*github.com/stretchr/testify/assert.Assertions).Equal`


## License

[MIT License][license-url]


[go-img]: https://img.shields.io/github/go-mod/go-version/maratori/pairedbrackets
[go-url]: /go.mod
[ci-img]: https://github.com/maratori/pairedbrackets/actions/workflows/ci.yml/badge.svg
[ci-url]: https://github.com/maratori/pairedbrackets/actions/workflows/ci.yml
[codecov-img]: https://codecov.io/gh/maratori/pairedbrackets/branch/main/graph/badge.svg?token=EGSPoXDeXP
[codecov-url]: https://codecov.io/gh/maratori/pairedbrackets
[codebeat-img]: https://codebeat.co/badges/650fdbf0-cad2-4533-979e-ee0e0f74edb8
[codebeat-url]: https://codebeat.co/projects/github-com-maratori-pairedbrackets-main
[codeclimate-img]: https://api.codeclimate.com/v1/badges/18392fd0a0ac261df437/maintainability
[codeclimate-url]: https://codeclimate.com/github/maratori/pairedbrackets/maintainability
[goreportcard-img]: https://goreportcard.com/badge/github.com/maratori/pairedbrackets
[goreportcard-url]: https://goreportcard.com/report/github.com/maratori/pairedbrackets
[license-img]: https://img.shields.io/github/license/maratori/pairedbrackets.svg
[license-url]: /LICENSE
[godoc-img]: https://pkg.go.dev/badge/github.com/maratori/pairedbrackets.svg
[godoc-url]: https://pkg.go.dev/github.com/maratori/pairedbrackets
