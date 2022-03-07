//go:build go1.18
// +build go1.18

package fuzzing

import (
	"strconv"
	"strings"
	"testing"

	fizzbuzz "github.com/caring/test/0-begin-here"
)

func FuzzFizzBuzz(f *testing.F) {
	f.Add(3, "Fizz")
	f.Add(5, "Buzz")
	f.Add(15, "FizzBuzz")
	f.Fuzz(func(t *testing.T, i int, s string) {
		defer func() {
			if r := recover(); r != nil {
				if i != 0 {
					t.Errorf("%q, nopanic", i)
				}
			}
		}()

		out, err := fizzbuzz.FizzBuzz(i)

		if i < 0 && err == nil {
			t.Errorf("%v, noerr", i)
		}

		if i > 0 && err != nil && out != "" {
			t.Errorf("%v, %v", out, err)
		}

		if i > 0 {
			if i%3 == 0 && !strings.HasPrefix(out, "fizz") {
				t.Errorf("%v, nofizz", i)
			}

			if i%5 == 0 && !strings.HasSuffix(out, "buzz") {
				t.Errorf("%v, nobuzz", i)
			}

			if i%15 == 0 && out != "fizzbuzz" {
				t.Errorf("%v, nofizzbuz", i)
			}

			if !(i%3 == 0 || i%5 == 0) && out != strconv.Itoa(i) {
				t.Errorf("%v, noint", i)
			}
		}
	})
}
