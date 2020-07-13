package kemba

import (
	"os"
	"testing"
)

func Test_New(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		k := New("test:kemba")
		if k.tag != "test:kemba" {
			t.Errorf("%#v, wanted %#v", k.tag, "test:kemba")
		}

		if k.allowed != "" {
			t.Errorf("%#v, wanted %#v", k.allowed, "")
		}
	})

	t.Run("simple w/ DEBUG set", func(t *testing.T) {
		os.Setenv("DEBUG", "test:*")

		k := New("test:kemba")
		if k.tag != "test:kemba" {
			t.Errorf("%#v, wanted %#v", k.tag, "test:kemba")
		}

		if k.allowed != "test:*" {
			t.Errorf("%#v, wanted %#v", k.allowed, "test:*")
		}

		os.Setenv("DEBUG", "")
	})
}
