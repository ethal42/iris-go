<div align="center">
  <h1>The Iris</h1>
  <i>A (power + color)ful library for Go.</i> <br/> <br/>
  
  [![License](https://img.shields.io/badge/license-MPL--2.0-lightblue)](LICENSE)
  [![Go Report Card](https://goreportcard.com/badge/github.com/ilmiofiume/iris-go)](https://goreportcard.com/report/github.com/ilmiofiume/iris-go)
  [![Go Reference](https://pkg.go.dev/badge/github.com/ilmiofiume/iris-go.svg)](https://pkg.go.dev/github.com/ilmiofiume/iris-go)
  [![Linux](https://github.com/ilmiofiume/iris-go/actions/workflows/linux.yml/badge.svg)](https://github.com/ilmiofiume/iris-go/actions/workflows/linux.yml)
  [![MacOS](https://github.com/ilmiofiume/iris-go/actions/workflows/macos.yml/badge.svg)](https://github.com/ilmiofiume/iris-go/actions/workflows/macos.yml)
  [![Windows](https://github.com/ilmiofiume/iris-go/actions/workflows/windows.yml/badge.svg)](https://github.com/ilmiofiume/iris-go/actions/workflows/windows.yml)
</div>

# How to Install?
Just use:
```
go get github.com/ilmiofiume/iris-go
```

# Example
```go
package main

import "github.com/ilmiofiume/iris-go"

func main() {
  c := iris.New(FgRed, BgWhite)
  c.Printf("Hello, %s\n!", "Iris")
  iris.Red("This is an error!!!!")
}
```

# License
Iris is licensed under the terms of [Mozilla Public License 2.0](LICENSE).
