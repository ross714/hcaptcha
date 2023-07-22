# Go-hCaptcha

[![go ref](https://pkg.go.dev/badge/github.com/ross714/hcaptcha.svg)](https://pkg.go.dev/github.com/ross714/hcaptcha)

Simple [hCaptcha](https://www.hcaptcha.com/) middleware written in Go, using fasthttp, which is up to 10x faster than net/http. 

This package works with any framework, since it only handles the verification.

Inspired by [flask-hcaptcha](https://github.com/KnugiHK/flask-hcaptcha) and [hcaptcha](https://github.com/kataras/hcaptcha).

## Installation
```sh
go get -u github.com/ross714/hcaptcha
```

## Getting Started
First of all, create an account on <https://www.hcaptcha.com/>, and attach a [new site](https://dashboard.hcaptcha.com/sites) for [development](https://docs.hcaptcha.com/#localdev). 

```go
import "github.com/ross714/hcaptcha"
```
```go
client := hcaptcha.New("secret_key", "site_key")
```
```go
res := client.Verify("token") // => Bool

if res {
  // Success
} else {
  // Failed
}
```

See the full example in [examples](examples) directory.

## License

This software is licensed under the [MIT License](LICENSE).
