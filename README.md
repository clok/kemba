# kemba
[![Go Report Card](https://goreportcard.com/badge/clok/kemba)](https://goreportcard.com/report/clok/kemba) [![Coverage Status](https://coveralls.io/repos/github/clok/kemba/badge.svg?branch=chore/test-coverage)](https://coveralls.io/github/clok/kemba?branch=chore/test-coverage) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/clok/kemba?tab=overview)

`debug` logging tool inspired by https://github.com/visionmedia/debug

#### Why is it named `kemba`?

`debug` is more generally considered to be [`runtime/debug`](https://golang.org/pkg/runtime/debug/) within Go. Since this takes heavy inspiration from my experiences using the `npm` module [`debug`](https://github.com/visionmedia/debug) I wanted to find a word that was somewhat connected to the inspiration. According to [Google translate](https://www.google.com/search?q=debug+in+icelandic) "debug" in English translated to Icelandic results in "kemba".

## Usage

The `kemba` logger reads the `DEBUG` and `KEMBA` environment variables to determine if a log line should be output. The logger outputs to `STDERR`.

When it is not set, the logger will immediately return, taking no action.

When the value is set (ex. `DEBUG=example:*,tool:details` and/or `KEMBA=plugin:fxn:start`), the logger will determine if it should be `enabled` when instantiated.

The value of these flags can be a simple regex alternative where a wildcard (`*`) are replaced with `.*` and all terms are prepended with `^` and appended with `$`. If a term does not include a wildcard, then an exact match it required.

To disabled colors, set the `NOCOLOR` environment variable to any value.

![example_output](https://user-images.githubusercontent.com/1429775/87724662-7112fd80-c781-11ea-86e7-95bd03c5c0a1.png)

```go
package main

import "github.com/clok/kemba"

type myType struct {
	a, b int
}

// When the DEBUG or KEMBA environment variable is set to DEBUG=example:* the kemba logger will output to STDERR
func main () {
    k := kemba.New("example:tag")
	
    var x = []myType{{1, 2}, {3, 4}}
    k.Printf("%#v", x)
    // Output to os.Stderr
    // example:tag []main.myType{main.myType{a:1, b:2}, main.myType{a:3, b:4}}

    k.Printf("%# v", x)
    k.Println(x)
    k.Log(x)
    // All result in the same output to os.Stderr
    // example:tag []main.myType{
    // example:tag     {a:1, b:2},
    // example:tag     {a:3, b:4},
    // example:tag }

    k1 := k.Extend("1")
    k1.Println("a string", 12, true)
    // Output to os.Stderr
    // example:tag:1 a string
    // example:tag:1 int(12)
    // example:tag: bool(true)
}
```

