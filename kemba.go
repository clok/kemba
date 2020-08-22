package kemba

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"github.com/kr/pretty"
	"hash/crc64"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// Kemba is a container struct for the Kemba logger.
// It used to manage the state of the logger.
// Currently all properties are not exported.
type Kemba struct {
	tag     string
	allowed string
	enabled bool
	logger  *log.Logger
	color   bool
	last   time.Time
}

var (
	table  = crc64.MakeTable(crc64.ISO)
	gs     = color.C256(uint8(240))
	colors = []int{
		20,
		21,
		26,
		27,
		32,
		33,
		38,
		39,
		40,
		41,
		42,
		43,
		44,
		45,
		56,
		57,
		62,
		63,
		68,
		69,
		74,
		75,
		76,
		77,
		78,
		79,
		80,
		81,
		92,
		93,
		98,
		99,
		112,
		113,
		128,
		129,
		134,
		135,
		148,
		149,
		160,
		161,
		162,
		163,
		164,
		165,
		166,
		167,
		168,
		169,
		170,
		171,
		172,
		173,
		178,
		179,
		184,
		185,
		196,
		197,
		198,
		199,
		200,
		201,
		202,
		203,
		204,
		205,
		206,
		207,
		208,
		209,
		214,
		215,
		220,
		221,
	}
)

// New Returns a Kemba logging instance. It will determine if the logger should
// bypass logging actions or be activated.
func New(tag string) *Kemba {
	allowed := getDebugFlagFromEnv()

	logger := Kemba{tag: tag, allowed: allowed}

	if logger.allowed != "" {
		logger.enabled = determineEnabled(tag, allowed)
		logger.color = os.Getenv("NOCOLOR") == ""
	} else {
		logger.enabled = false
		logger.color = false
	}

	var prefix string
	if logger.enabled {
		if logger.color {
			s := PickColor(tag)
			prefix = s.Sprintf("%s ", tag)
		} else {
			prefix = fmt.Sprintf("%s ", tag)
		}

		logger.logger = log.New(os.Stderr, prefix, log.Lmsgprefix)
		logger.last = time.Now()
	}

	return &logger
}

// Printf is a convenience wrapper that will apply pretty.Formatter to the passed in variables.
//
// Calling Printf(f, x, y) is equivalent to fmt.Printf(f, pretty.Formatter(x), pretty.Formatter(y)).
func (k *Kemba) Printf(format string, v ...interface{}) {
	if k.enabled {
		elapsed := k.determineElapsed()

		var buf bytes.Buffer
		_, _ = pretty.Fprintf(&buf, format, v...)

		showDelta := true
		k.printBuffer(buf, elapsed, &showDelta)
	}
}

// Println is a convenience wrapper that will apply pretty.Formatter to the passed in variables.
//
// Calling Println(x, y) is equivalent to fmt.Println(pretty.Formatter(x), pretty.Formatter(y)),
// but each operand is formatted with "%# v".
func (k *Kemba) Println(v ...interface{}) {
	if k.enabled {
		showDelta := true
		for _, x := range v {
			elapsed := k.determineElapsed()

			var buf bytes.Buffer
			_, _ = pretty.Fprintf(&buf, "%# v", x)

			k.printBuffer(buf, elapsed, &showDelta)
		}
	}
}

// Log is an alias to Println
func (k *Kemba) Log(v ...interface{}) {
	k.Println(v...)
}

// Extend returns a new Kemba logger instance that has appended the provided tag to the original logger.
//
// New logger instance will have original `tag` value delimited with a `:` and appended with the new extended `tag` input.
//
// Example:
//     k := New("test:original)
//     k.Log("test")
//     ke := k.Extend("plugin")
//     ke.Log("test extended")
//
// Output:
//     test:original test
//     test:original:plugin test extended
func (k *Kemba) Extend(tag string) *Kemba {
	exTag := fmt.Sprintf("%s:%s", k.tag, tag)
	return New(exTag)
}

// PickColor will return the same color based on input string.
//
// We want to pick the same color for a given tag to ensure consistent output behavior.
func PickColor(tag string) *color.Color256 {
	// Generate an 8 byte checksum to pass into Rand.seed
	seed := crc64.Checksum([]byte(tag), table)
	rand.Seed(int64(seed))
	v := rand.Intn(len(colors) - 1)
	s := color.C256(uint8(colors[v]))
	return &s
}

// printBuffer will append the elapsed time delta to the first line of the provided buffer
// if the showDelta parameter is true. Otherwise, this method prints the buffer lines to STDERR
func (k *Kemba) printBuffer(buf bytes.Buffer, elapsed time.Duration, showDelta *bool) {
	s := bufio.NewScanner(&buf)
	for s.Scan() {
		if *showDelta {
			var ft string
			if k.color {
				ft = gs.Sprintf("+%s", elapsed.Truncate(time.Millisecond))
			} else {
				ft = fmt.Sprintf("+%s", elapsed.Truncate(time.Millisecond))
			}
			k.logger.Printf("%s %s\n", s.Text(), ft)
			*showDelta = false
		} else {
			k.logger.Print(s.Text())
		}
	}
}

// getDebugFlagFromEnv considers both the value of DEBUG and KEMBA env values
// to determine the resulting logging flags to pass to the loggers.
func getDebugFlagFromEnv() string {
	dEnv := os.Getenv("DEBUG")
	kEnv := os.Getenv("KEMBA")

	var s []string
	if dEnv != "" {
		s = append(s, dEnv)
	}
	if kEnv != "" {
		s = append(s, kEnv)
	}

	if len(s) > 1 {
		return strings.Join(s, ",")
	} else if len(s) > 0 {
		return s[0]
	}

	return ""
}

// determineElapsed will determine the time delta from between the last log event for this
// Kemba logger and return the elapsed time.
func (k *Kemba) determineElapsed() time.Duration {
	now := time.Now()
	elapsed := now.Sub(k.last)
	k.last = now

	return elapsed
}

// determineEnabled will check the value of DEBUG and KEMBA environment variables to generate regex to test against the tag
//
// If no * in string, then assume exact match
// Else
// It will split by , and perform
// It will, replace * with .*
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
		} else if !a {
			a = l == tag
		}
	}
	return a
}
