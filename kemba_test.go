package kemba

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_New(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		k := New("test:kemba")
		if k.enabled != false {
			t.Error("Logger should be disabled")
		}

		if k.tag != "test:kemba" {
			t.Errorf("%#v, wanted %#v", k.tag, "test:kemba")
		}

		if k.allowed != "" {
			t.Errorf("%#v, wanted %#v", k.allowed, "")
		}
	})

	t.Run("simple w/ DEBUG set", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		k := New("test:kemba")
		if k.enabled != true {
			t.Error("Logger should be enabled")
		}

		if k.tag != "test:kemba" {
			t.Errorf("%#v, wanted %#v", k.tag, "test:kemba")
		}

		if k.allowed != "test:*" {
			t.Errorf("%#v, wanted %#v", k.allowed, "test:*")
		}

		_ = os.Setenv("DEBUG", "")
	})

	t.Run("exact match tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:kemba")

		k := New("test:kemba")
		if k.enabled != true {
			t.Error("Logger should be enabled")
		}
	})

	t.Run("exact match tag [miss]", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:kemba:1")

		k := New("test:kemba")
		if k.enabled != false {
			t.Error("Logger should be disabled")
		}
	})

	t.Run("fuzzy match tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "test:*")

		k := New("test:kemba")
		if k.enabled != true {
			t.Error("Logger should be enabled")
		}
	})

	t.Run("fuzzy match tag", func(t *testing.T) {
		_ = os.Setenv("DEBUG", "*kemba")

		k := New("test:kemba:fail")
		if k.enabled != false {
			t.Error("Logger should be disabled")
		}
	})
}

func Example() {
	_ = os.Setenv("DEBUG", "example:*")
	k := New("example:tag")
	k1 := New("example:tag:1")

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

	k1.Println("a string", 12, true)
	// Output to os.Stderr
	// example:tag:1 a string
	// example:tag:1 int(12)
	// example:tag: bool(true)
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

	k1 := New("test:kemba:1")
	k1.Printf("%s", "Hello 1")

	k2 := New("test:kemba:2")
	k2.Printf("%s", "Hello 2")

	k3 := New("test:kemba:3")
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

		if string(out) != "" {
			t.Errorf("%#v, wanted %#v", string(out), "")
		}
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

		wantMsg := "test:kemba key: test value: 1337\n"
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
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
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
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

		wantMsg := "test:kemba []kemba.myType{kemba.myType{a:1, b:2}, kemba.myType{a:3, b:4}, kemba.myType{a:5, b:6}}\n"
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
		}

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
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
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

		if !strings.Contains(string(out), "key: test value: 1337") {
			t.Errorf("Expected string %#v to contain %#v", string(out), "key: test value: 1337")
		}
		if !strings.Contains(string(out), "test:kemba") {
			t.Errorf("Expected string %#v to contain %#v", string(out), "test:kemba")
		}
		if os.Getenv("CI") == "" {
			if !strings.Contains(string(out), "\x1b[") {
				t.Errorf("Expected string %#v to contain %#v", string(out), "\x1b[")
			}
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
	t.Run("should do nothing when no DEBUG flag is set", func(t *testing.T) {
		rescueStderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w

		k := New("test:kemba")
		k.Println("test")

		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr

		if string(out) != "" {
			t.Errorf("%#v, wanted %#v", string(out), "")
		}
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

		wantMsg := "test:kemba test\n"
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
		}

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
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
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

		wantMsg := `test:kemba this
test:kemba is
test:kemba a
test:kemba multiline
test:kemba string
`
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
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

		wantMsg := `test:kemba []kemba.myType{
test:kemba     {a:1, b:2},
test:kemba     {a:3, b:4},
test:kemba     {a:5, b:6},
test:kemba }
`
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
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

		if !strings.Contains(string(out), "int(1337)") {
			t.Errorf("Expected string %#v to contain %#v", string(out), "int(1337)")
		}
		if !strings.Contains(string(out), "test:kemba") {
			t.Errorf("Expected string %#v to contain %#v", string(out), "test:kemba")
		}
		if os.Getenv("CI") == "" {
			if !strings.Contains(string(out), "\x1b[") {
				t.Errorf("Expected string %#v to contain %#v", string(out), "\x1b[")
			}
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
		if string(out) != wantMsg {
			t.Errorf("%#v, wanted %#v", string(out), wantMsg)
		}

		_ = os.Setenv("DEBUG", "")
		_ = os.Setenv("NOCOLOR", "")
	})
}
