package kemba

import (
	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
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
		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k1 := New("test:kemba")
		k2 := New("test:kemba")
		k1.Log("this shoulc be the same color")
		k2.Log("this shoulc be the same color")

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		lines := strings.Split(string(out), "\n")

		is.Equal(lines[0], lines[1], "Both lines should have the same color prompt")

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
		_ = os.Setenv("DEBUG", "*kemba")

		k := New("test:kemba:fail")
		is.False(k.enabled, "Logger should NOT be enabled")
	})
}

func Example() {
	_ = os.Setenv("DEBUG", "example:*")
	k := New("example:tag")

	type myType struct {
		a, b int
	}

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

	// Create a new extended logger with a new tag
	k1 := k.Extend("1")
	k1.Println("a string", 12, true)
	// Output to os.Stderr
	// example:tag:1 a string
	// example:tag:1 int(12)
	// example:tag:1 bool(true)
	_ = os.Setenv("DEBUG", "")

	// Output:
	// example:tag []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}}
	// example:tag []kemba.myType{
	// example:tag     {a:1, b:2},
	// example:tag     {a:3, b:4},
	// example:tag }
	// example:tag []kemba.myType{
	// example:tag     {a:1, b:2},
	// example:tag     {a:3, b:4},
	// example:tag }
	// example:tag []kemba.myType{
	// example:tag     {a:1, b:2},
	// example:tag     {a:3, b:4},
	// example:tag }
	// example:tag:1 a string
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
	// test:kemba Hello
	// test:kemba:1 Hello 1
	// test:kemba:2 Hello 2
	// test:kemba:3 Hello 3
	// test:kemba:2 []string{"test", "again", "third"}
	// test:kemba:1 map[string]int{"test":1, "again":1337, "third":732}
	// test:kemba:3 []kemba.myType{
	// test:kemba:3     {a:1, b:2},
	// test:kemba:3     {a:3, b:4},
	// test:kemba:3     {a:5, b:6},
	// test:kemba:3 }
	// test:kemba:2 []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}, kemba.myType{a:5, b:6}}
	// test:kemba map[string]int{"again":1337, "test":1, "third":732} []string{"test", "again", "third"} map[string]int{"again":1337, "test":1, "third":732} []string{"test", "again", "third"} map[string]int{"again":1337, "test":1, "third":732} []string{"test", "again", "third"}
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
	// test:kemba Hello
	// test:kemba []kemba.myType{
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
	// test:kemba []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}, kemba.myType{a:5, b:6}}
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

		is.Equal("test:kemba key: test value: 1337\n", string(out))

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

		wantMsg := `test:kemba this
test:kemba is
test:kemba a
test:kemba multiline
test:kemba string
`
		is.Equal(wantMsg, string(out))

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

		wantMsg := "test:kemba []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}, kemba.myType{a:5, b:6}}\n"
		is.Equal(wantMsg, string(out))

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

		wantMsg := `test:kemba []kemba.myType{
test:kemba     {a:1, b:2},
test:kemba     {a:3, b:4},
test:kemba     {a:5, b:6},
test:kemba }
`
		is.Equal(wantMsg, string(out))

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
	// test:kemba Hello
	// test:kemba []kemba.myType{
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

		is.Equal("test:kemba test\n", string(out))

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

		wantMsg := `test:kemba test
test:kemba int(1337)
`
		is.Equal(wantMsg, string(out))

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

		wantMsg := `test:kemba this
test:kemba is
test:kemba a
test:kemba multiline
test:kemba string
`
		is.Equal(wantMsg, string(out))

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

		wantMsg := `test:kemba []kemba.myType{
test:kemba     {a:1, b:2},
test:kemba     {a:3, b:4},
test:kemba     {a:5, b:6},
test:kemba }
`
		is.Equal(wantMsg, string(out))

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
	// test:kemba Hello
	// test:kemba []kemba.myType{
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

		wantMsg := `test:kemba []kemba.myType{
test:kemba     {a:1, b:2},
test:kemba     {a:3, b:4},
test:kemba     {a:5, b:6},
test:kemba }
`
		is.Equal(wantMsg, string(out))

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

		is.Equal("test:kemba:extended-walrus key: test value: 1337\n", string(out))

		_ = os.Setenv("DEBUG", "")
	})
}

func Test_Private_PickColor(t *testing.T) {
	is := assert.New(t)

	t.Run("should return the same color for a given string", func(t *testing.T) {
		out := pickColor("test:kemba")

		c := color.Color256{81, 0}
		is.Equal(c.Value(), out.Value())
	})
}
