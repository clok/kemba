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
	"regexp"
	"strings"
	"time"
)

type kLog struct {
	tag     string
	allowed string
	enabled bool
	logger  *log.Logger
	color   bool
}

// New Returns a kLog logging instance
func New(tag string) *kLog {
	allowed := os.Getenv("DEBUG")

	logger := kLog{tag: tag, allowed: allowed}

	if os.Getenv("DEBUG") != "" {
		logger.enabled = determineEnabled(tag, allowed)
		logger.color = os.Getenv("NOCOLOR") == ""
	} else {
		logger.enabled = false
		logger.color = false
	}

	var prefix string
	if logger.enabled {
		if logger.color {
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

// Printf is a convenience wrapper that will apply pretty.Formatter to the passed in variables.
// Calling Printf(f, x, y) is equivalent to fmt.Printf(f, Formatter(x), Formatter(y)).
func (k kLog) Printf(format string, v ...interface{}) {
	if k.enabled {
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
func (k kLog) Println(v ...interface{}) {
	if k.enabled {
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
func (k kLog) Log(v ...interface{}) {
	k.Println(v...)
}

// determineEnabled will check the value of DEBUG environment variable to generate regex to test against the tag
// It will split by , and perform
// It will, replace * with .*
// If no * then exact match
func determineEnabled(tag string, allowed string) bool {
	var a bool
	for _, l := range strings.Split(allowed, ",") {
		if strings.Contains(l, "*") {
			reg := strings.ReplaceAll(l, "*", ".*")
			if !strings.HasPrefix(reg, "^") {
				reg = fmt.Sprintf("^%s", reg)
			}

			if !strings.HasSuffix(reg, "$") {
				reg = fmt.Sprintf("%s$", reg)
			}

			if !a {
				a, _ = regexp.Match(reg, []byte(tag))
			}
		} else {
			if !a {
				a = l == tag
			}
		}
	}
	return a
}

// toggleColor with turn color on and off.
// TODO: enable functionality
func (k kLog) toggleColor() {
	k.color = !k.color
}
