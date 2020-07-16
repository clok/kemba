package kemba

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kr/pretty"
	"gopkg.in/gookit/color.v1"
	"log"
	"math/rand"
	"os"
	"time"
)

type Kemba struct {
	tag     string
	filter  interface{}
	allowed string
	enabled bool
	logger  *log.Logger
	color   bool
}

// New Returns a Kemba logging instance
func New(tag string) *Kemba {
	allowed := os.Getenv("DEBUG")

	logger := Kemba{tag: tag, allowed: allowed}

	if allowed != "" {
		logger.enabled = true
		logger.color = true
	} else {
		logger.enabled = false
		logger.color = false
	}

	var prefix string
	if logger.enabled {
		if os.Getenv("NOCOLOR") == "" && logger.color {
			rand.Seed(time.Now().UnixNano())
			cint := rand.Intn(230) + 1
			if cint == 8 {
				cint += 1
			}
			if cint == 16 {
				cint += 1
			}
			s := color.C256(uint8(cint))
			prefix = s.Sprintf("%s ", tag)
		} else {
			prefix = fmt.Sprintf("%s ", tag)
		}

		logger.logger = log.New(os.Stderr, prefix, log.Lmsgprefix)
	}

	return &logger
}

// toggleColor with turn color on and off.
// TODO: enable functionality
func (k Kemba) toggleColor() {
	k.color = !k.color
}

// Printf is a convenience wrapper that will apply pretty.Formatter to the passed in variables.
// Calling Printf(f, x, y) is equivalent to fmt.Printf(f, Formatter(x), Formatter(y)).
func (k Kemba) Printf(format string, v ...interface{}) {
	if k.enabled {
		// TODO: add in regex/lookup table
		if k.allowed == "" {
			return
		}

		var buf bytes.Buffer
		_, _ = pretty.Fprintf(&buf, format, v...)

		s := bufio.NewScanner(&buf)
		for s.Scan() {
			k.logger.Print(s.Text())
		}
	}
}

// Println is a convenience wrapper that will apply pretty.Formatter to the passed in variables.
// Calling Println(x, y) is equivalent to fmt.Println(Formatter(x), Formatter(y)), but each operand is formatted with "%# v".
func (k Kemba) Println(v ...interface{}) {
	if k.enabled {
		// TODO: add in regex/lookup table
		if k.allowed == "" {
			return
		}

		for _, x := range v {
			var buf bytes.Buffer
			_, _ = pretty.Fprintf(&buf, "%# v", x)

			s := bufio.NewScanner(&buf)
			for s.Scan() {
				k.logger.Print(s.Text())
			}
		}
	}
}

// Log is an alias to Println
func (k Kemba) Log(v ...interface{}) {
	k.Println(v...)
}
