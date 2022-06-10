//nolint:structcheck
package kemba

import (
	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_New(t *testing.T) {
	is := assert.New(t)

	t.Run("simple", func(t *testing.T) {
		k := New("test:kemba")
		is.False(k.enabled, "Logger should NOT be enabled")
		is.Equal("test:kemba", k.tag)
		is.Equal("", k.allowed)
	})

	t.Run("simple w/ DEBUG set", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		k := New("test:kemba")
		is.True(k.enabled, "Logger should be enabled")
		is.Equal("test:kemba", k.tag)
		is.Equal("test:*", k.allowed)

		_ = os.Setenv("DEBUG", "")
	})

	t.Run("should pick the same color for a given tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		k1 := New("test:kemba")
		k2 := New("test:kemba")

		is.Equal(k1.color, k2.color)

		_ = os.Setenv("DEBUG", "")
	})

	t.Run("exact match tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:kemba")

		k := New("test:kemba")
		is.True(k.enabled, "Logger should be enabled")
	})

	t.Run("exact match tag [miss]", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:kemba:1")

		k := New("test:kemba")
		is.False(k.enabled, "Logger should NOT be enabled")
	})

	t.Run("fuzzy match tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		k := New("test:kemba")
		is.True(k.enabled, "Logger should be enabled")
	})

	t.Run("fuzzy match tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "*kemba*")

		k := New("test:kemba:fail")
		is.True(k.enabled, "Logger should be enabled")
	})

	t.Run("fuzzy match tag [failure]", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "*kemba")

		k := New("test:kemba:fail")
		is.False(k.enabled, "Logger should NOT be enabled")
	})

	t.Run("fuzzy match tag [mid star]", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*:fail")

		k := New("test:kemba:fail")
		is.True(k.enabled, "Logger should be enabled")
	})
}

func Example() {
	_ = os.Setenv("DEBUG", "example:*")
	// OR
	// _ = os.Setenv("KEMBA", "example:*")
	k := New("example:tag")

	type myType struct {
		a, b int
	}

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
	// example:tag []main.myType{ +XXs
	// example:tag     {a:1, b:2},
	// example:tag     {a:3, b:4},
	// example:tag }

	// Create a new extended logger with a new tag
	k1 := k.Extend("1")
	k1.Println("a string", 12, true)
	// Output to os.Stderr
	// example:tag:1 a string +0s
	// example:tag:1 int(12)
	// example:tag:1 bool(true)
	_ = os.Setenv("DEBUG", "")

	// Output:
	// example:tag []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}} +0s
	// example:tag []kemba.myType{ +250ms
	// example:tag     {a:1, b:2},
	// example:tag     {a:3, b:4},
	// example:tag }
	// example:tag []kemba.myType{ +0s
	// example:tag     {a:1, b:2},
	// example:tag     {a:3, b:4},
	// example:tag }
	// example:tag []kemba.myType{ +105ms
	// example:tag     {a:1, b:2},
	// example:tag     {a:3, b:4},
	// example:tag }
	// example:tag:1 a string +0s
	// example:tag:1 int(12)
	// example:tag:1 bool(true)
}

func ExampleKemba_Printf() {
	_ = os.Setenv("DEBUG", "test:*")
	k := New("test:kemba")
	k.Printf("%s", "Hello")

	k1 := k.Extend("1")
	k1.Printf("%s", "Hello 1")

	k2 := k.Extend("2")
	k2.Printf("%s", "Hello 2")

	k3 := k.Extend("3")
	k3.Printf("%s", "Hello 3")

	s := []string{"test", "again", "third"}
	k2.Printf("%# v", s)

	m := map[string]int{
		"test":  1,
		"again": 1337,
		"third": 732,
	}
	k1.Printf("%# v", m)

	type myType struct {
		a int
		b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}
	k3.Printf("%# v", x)
	k2.Printf("%#v", x)

	k.Printf("%#v %#v %#v %#v %#v %#v", m, s, m, s, m, s)
	_ = os.Setenv("DEBUG", "")

	// Output:
	// test:kemba Hello +0s
	// test:kemba:1 Hello 1 +0s
	// test:kemba:2 Hello 2 +0s
	// test:kemba:3 Hello 3 +0s
	// test:kemba:2 []string{"test", "again", "third"} +0s
	// test:kemba:1 map[string]int{"again":1337, "third":732, "test":1} +0s
	// test:kemba:3 []kemba.myType{ +0s
	// test:kemba:3     {a:1, b:2},
	// test:kemba:3     {a:3, b:4},
	// test:kemba:3     {a:5, b:6},
	// test:kemba:3 }
	// test:kemba:2 []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}, kemba.myType{a:5, b:6}} +0s
	// test:kemba map[string]int{"again":1337, "test":1, "third":732} []string{"test", "again", "third"} map[string]int{"again":1337, "test":1, "third":732} []string{"test", "again", "third"} map[string]int{"again":1337, "test":1, "third":732} []string{"test", "again", "third"} +0s
}

func ExampleKemba_Printf_expanded() {
	_ = os.Setenv("DEBUG", "test:*")
	k := New("test:kemba")
	k.Printf("%s", "Hello")

	type myType struct {
		a int
		b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}

	// NOTE: The "%# v" operand for the Printf format.
	k.Printf("%# v", x)
	_ = os.Setenv("DEBUG", "")

	// Output:
	// test:kemba Hello +0s
	// test:kemba []kemba.myType{ +0s
	// test:kemba     {a:1, b:2},
	// test:kemba     {a:3, b:4},
	// test:kemba     {a:5, b:6},
	// test:kemba }
}

func ExampleKemba_Printf_compact() {
	_ = os.Setenv("DEBUG", "test:*")
	k := New("test:kemba")

	type myType struct {
		a int
		b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}

	// NOTE: The "%#v" operand for the Printf format.
	k.Printf("%#v", x)
	_ = os.Setenv("DEBUG", "")

	// Output:
	// test:kemba []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}, kemba.myType{a:5, b:6}} +0s
}

func Test_Printf(t *testing.T) {
	is := assert.New(t)

	t.Run("should do nothing when no DEBUG flag is set", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "")
		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Printf("key: %s value: %d", "test", 1337)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Equal("", string(out))
	})

	t.Run("should prepend tag on simple string", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Printf("key: %s value: %d", "test", 1337)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Regexp(`^test:kemba key: test value: 1337 \+\d+\S+\n$`, string(out))

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on multiline string", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		s := `this
is
a
multiline
string`
		k.Printf("%s", s)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		for i, ln := range strings.Split(string(out), "\n") {
			if ln == "" {
				continue
			}
			if i == 0 {
				is.Regexp(`^test:kemba\sthis\s\+\d+\S{1,2}`, ln)
			} else {
				is.Regexp(`^test:kemba\s\w+$`, ln)
			}
		}

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on typed struct", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		type myType struct {
			a int
			b int
		}
		var x = []myType{{1, 2}, {3, 4}, {5, 6}}
		k.Printf("%#v", x)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Regexp(`^test:kemba \[\]kemba.myType\{kemba.myType\{a:1, b:2\}, kemba.myType\{a:3, b:4\}, kemba.myType\{a:5, b:6\}\} \+\d+\w\n`, string(out))

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on typed struct (multiline)", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		type myType struct {
			a, b int
		}
		var x = []myType{{1, 2}, {3, 4}, {5, 6}}
		k.Printf("%# v", x)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		for i, ln := range strings.Split(string(out), "\n") {
			if ln == "" {
				continue
			}
			if i == 0 {
				is.Regexp(`^test:kemba\s\[\]kemba\.myType\{\s\+\d+\S{1,2}$`, ln)
			} else {
				is.Regexp(`^test:kemba\s+[\{\}\:ab123456,\s]+$`, ln)
			}
		}

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on simple string w/ color", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Printf("key: %s value: %d", "test", 1337)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Contains(string(out), "key: test value: 1337")
		is.Contains(string(out), "test:kemba")
		if os.Getenv("CI") == "" {
			is.Contains(string(out), "\x1b[")
		}

		_ = os.Setenv("DEBUG", "")
	})
}

func ExampleKemba_Println() {
	_ = os.Setenv("DEBUG", "test:*")
	k := New("test:kemba")
	k.Printf("%s", "Hello")

	type myType struct {
		a int
		b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}
	k.Println(x)

	_ = os.Setenv("DEBUG", "")

	// Output:
	// test:kemba Hello +0s
	// test:kemba []kemba.myType{ +0s
	// test:kemba     {a:1, b:2},
	// test:kemba     {a:3, b:4},
	// test:kemba     {a:5, b:6},
	// test:kemba }
}

func Test_Println(t *testing.T) {
	is := assert.New(t)

	t.Run("should do nothing when no DEBUG flag is set", func(t *testing.T) {
		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Println("test")

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Equal("", string(out))
	})

	t.Run("should prepend tag on simple string", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Println("test")

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Regexp(`^test:kemba test \+\d+\S+\n$`, string(out))

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on lines for each variable passed", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Println("test", 1337)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		for i, ln := range strings.Split(string(out), "\n") {
			switch i {
			case 0:
				is.Regexp(`^test:kemba\s.*.\s\+\d+\S{1,2}`, ln)
			case 1:
				is.Regexp(`^test:kemba\sint\(1337\)$`, ln)
			default:
				continue
			}
		}

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on multiline string", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		s := `this
is
a
multiline
string`
		k.Println(s)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		for i, ln := range strings.Split(string(out), "\n") {
			switch i {
			case 0:
				is.Regexp(`^test:kemba\s\w+.\s\+\d+\S{1,2}$`, ln)
			case 1, 2, 3, 4:
				is.Regexp(`^test:kemba\s\w+$`, ln)
			default:
				continue
			}
		}

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on typed struct", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		type myType struct {
			a int
			b int
		}
		var x = []myType{{1, 2}, {3, 4}, {5, 6}}
		k.Println(x)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		for i, ln := range strings.Split(string(out), "\n") {
			if ln == "" {
				continue
			}
			if i == 0 {
				is.Regexp(`^test:kemba\s\[\]kemba\.myType\{\s\+\d+\S{1,2}$`, ln)
			} else {
				is.Regexp(`^test:kemba\s+[\{\}\:ab123456,\s]+$`, ln)
			}
		}

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})

	t.Run("should prepend tag on simple string w/ color", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Println(1337)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Contains(string(out), "int(1337)")
		is.Contains(string(out), "test:kemba")
		if os.Getenv("CI") == "" {
			is.Contains(string(out), "\x1b[")
		}

		_ = os.Setenv("DEBUG", "")
	})

}

func ExampleKemba_Log() {
	_ = os.Setenv("DEBUG", "test:*")
	k := New("test:kemba")
	k.Printf("%s", "Hello")

	type myType struct {
		a, b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}
	k.Log(x)

	_ = os.Setenv("DEBUG", "")

	// Output:
	// test:kemba Hello +0s
	// test:kemba []kemba.myType{ +0s
	// test:kemba     {a:1, b:2},
	// test:kemba     {a:3, b:4},
	// test:kemba     {a:5, b:6},
	// test:kemba }
}

func Test_Log(t *testing.T) {
	is := assert.New(t)

	t.Run("should prepend tag on typed struct", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		type myType struct {
			a, b int
		}
		var x = []myType{{1, 2}, {3, 4}, {5, 6}}
		k.Log(x)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		for i, ln := range strings.Split(string(out), "\n") {
			if ln == "" {
				continue
			}
			if i == 0 {
				is.Regexp(`^test:kemba\s\[\]kemba\.myType\{\s\+\d+\S{1,2}$`, ln)
			} else {
				is.Regexp(`^test:kemba\s+[\{\}\:ab123456,\s]+$`, ln)
			}
		}

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})
}

func Test_Extend(t *testing.T) {
	is := assert.New(t)

	t.Run("should extend the original tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")
		_ = os.Setenv("NOCOLOR", "1")

		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		ke := k.Extend("extended-walrus")
		ke.Printf("key: %s value: %d", "test", 1337)

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		is.Regexp(`^test:kemba:extended-walrus key: test value: 1337 \+\d+\S+\n$`, string(out))

		_ = os.Setenv("DEBUG", "")
	})
}

func Test_PickColor(t *testing.T) {
	is := assert.New(t)

	t.Run("should return the same color for a given string", func(t *testing.T) {
		out := PickColor("test:kemba")

		c := color.Color256{81, 0}
		is.Equal(c.Value(), out.Value())
	})
}

func Test_Private_getDebugFlagFromEnv(t *testing.T) {
	is := assert.New(t)

	t.Run("should return an empty string", func(t *testing.T) {
		out := getDebugFlagFromEnv()

		is.Equal("", out)
	})

	t.Run("should return value of DEBUG", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		out := getDebugFlagFromEnv()

		is.Equal("test:*", out)

		_ = os.Setenv("DEBUG", "")
	})

	t.Run("should return value of KEMBA", func(t *testing.T) {
		_ = os.Setenv("KEMBA", "test:*")

		out := getDebugFlagFromEnv()

		is.Equal("test:*", out)

		_ = os.Setenv("DEBUG", "")
	})

	t.Run("should return value of DEBUG appended with KEMBA", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "debug:*")
		_ = os.Setenv("KEMBA", "kemba:*")

		out := getDebugFlagFromEnv()

		is.Equal("debug:*,kemba:*", out)

		_ = os.Setenv("KEMBA", "")
		_ = os.Setenv("KEMBA", "")
	})
}
