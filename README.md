# pairedbrackets <br> [![go.mod version][go-img]][go-url] [![CI][ci-img]][ci-url] [![Codecov][codecov-img]][codecov-url] [![Codebeat][codebeat-img]][codebeat-url] [![Maintainability][codeclimate-img]][codeclimate-url] [![Go Report Card][goreportcard-img]][goreportcard-url] [![License][license-img]][license-url] [![Go Reference][godoc-img]][godoc-url]

Linter ensures that notation rule from [this article](https://www.yegor256.com/2014/10/23/paired-brackets-notation.html) is respected:
"a bracket should either start/end a line or be paired on the same line".

```go
http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
	...
})
```
```go
http.HandleFunc(
	"/api",
	func(w http.ResponseWriter, r *http.Request) {
		...
	},
)
```

## Examples

### Function parameters

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func DoSomething(a int,
	b string, c bool) {
	...
}
```
```go
func DoSomething(
	a int,
	b string,
	c bool) {
	...
}
```
```go
func DoSomething(a int,
	b string,
	c bool,
) {
	...
}
```

</td><td>

```go
func DoSomething(a int, b string, c bool) {
	...
}
```
```go
func DoSomething(
	a int,
	b string,
	c bool,
) {
	...
}
```
```go
func DoSomething(
	a int, b string, c bool,
) {
	...
}
```
```go
func DoSomething(
	a int, b string,
	c bool,
) {
	...
}
```

</td></tr>

<tr><td>

```go
func DoSomething(a int, b string, c struct{
	X int
	Y string
},
) {
	...
}
```

</td><td>

```go
func DoSomething(a int, b string, c struct{
	X int
	Y string
}) {
	...
}
```

</td></tr>
</tbody></table>

### Function type parameters

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func DoSomething[T int,
	V string]() {
	...
}
```
```go
func DoSomething[
	T int,
	V string]() {
	...
}
```

</td><td>

```go
func DoSomething[T int, V string]() {
	...
}
```
```go
func DoSomething[
	T int,
	V string,
]() {
	...
}
```

</td></tr>
</tbody></table>

### Function returns

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func DoSomething() (int,
	error) {
	...
}
```

```go
func DoSomething() (
	int,
	error) {
	...
}
```

</td><td>

```go
func DoSomething() (int, error) {
	...
}
```
```go
func DoSomething() (
	int,
	error,
) {
	...
}
```

</td></tr>
</tbody></table>


## Usage as standalone linter

### Install

```shell
go install github.com/maratori/pairedbrackets@latest
```

### Run

```shell
pairedbrackets ./...
```


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
