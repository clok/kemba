# kemba
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://github.com/clok/kemba/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/clok/kemba)](https://goreportcard.com/report/clok/kemba)
[![Coverage Status](https://coveralls.io/repos/github/clok/kemba/badge.svg)](https://coveralls.io/github/clok/kemba)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/clok/kemba?tab=overview)
[![Mentioned in Awesome
Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)

`debug` logging tool inspired by https://github.com/visionmedia/debug

#### Why is it named `kemba`?

`debug` is more generally considered to be [`runtime/debug`](https://golang.org/pkg/runtime/debug/) within Go. Since this takes heavy inspiration from my experiences using the `npm` module [`debug`](https://github.com/visionmedia/debug) I wanted to find a word that was somewhat connected to the inspiration. According to [Google translate](https://www.google.com/search?q=debug+in+icelandic) "debug" in English translated to Icelandic results in "kemba".

## Usage

The `kemba` logger reads the `DEBUG` and `KEMBA` environment variables to determine if a log line should be output. The logger outputs to `STDERR`.

When it is not set, the logger will immediately return, taking no action.

When the value is set (ex. `DEBUG=example:*,tool:details` and/or `KEMBA=plugin:fxn:start`), the logger will determine if it should be `enabled` when instantiated.

The value of these flags can be a simple regex alternative where a wildcard (`*`) are replaced with `.*` and all terms are prepended with `^` and appended with `$`. If a term does not include a wildcard, then an exact match it required.

To disabled colors, set the `NOCOLOR` environment variable to any value.

![image](https://user-images.githubusercontent.com/1429775/88557149-7973ff80-cfef-11ea-8ec2-ff332fd1b25f.png)

```go
package main

import (
    "time"

	"github.com/clok/kemba"
)

type myType struct {
	a, b int
}

// When the DEBUG or KEMBA environment variable is set to DEBUG=example:* the kemba logger will output to STDERR
func main () {
    k := kemba.New("example:tag")
	
    var x = []myType{{1, 2}, {3, 4}}
    k.Printf("%#v", x)
    // Output to os.Stderr
    // example:tag []main.myType{main.myType{a:1, b:2}, main.myType{a:3, b:4}} +0s

    // Artificial delay to demonstrate the time tagging
    time.Sleep(250 * time.Millisecond)
    k.Printf("%# v", x)
    k.Println(x)

    // Artificial delay to demonstrate the time tagging
    time.Sleep(100 * time.Millisecond)
    k.Log(x)
    // All result in the same output to os.Stderr
    // example:tag []main.myType{ +XXms
    // example:tag     {a:1, b:2},
    // example:tag     {a:3, b:4},
    // example:tag }

    // Create a new logger with an extended tag
    k1 := k.Extend("1")
    k1.Println("a string", 12, true)
    // Output to os.Stderr
    // example:tag:1 a string +0s
    // example:tag:1 int(12)
    // example:tag:1 bool(true)
}
```

