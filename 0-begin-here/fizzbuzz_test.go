package fizzbuzz_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/matryer/is"

	fizzbuzz "github.com/caring/test/0-begin-here"
)

// ExampleFizzBuzz is a simple test to demonstrate how to use a function.
// it will compare what is printed against what is shown in the Output comment below.
func ExampleFizzBuzz() {
	print(fizzbuzz.FizzBuzz(3))
	print(fizzbuzz.FizzBuzz(5))
	print(fizzbuzz.FizzBuzz(15))

	// Output:
	// fizz
	// buzz
	// fizzbuzz
}
func print(s string, err error) {
	fmt.Println(s)
}

// TestFizzBuzz_Simple shows how to use the different assertions when testing.
func TestFizzBuzz_Simple(t *testing.T) { 
	// is has a minimal set of functions that can be used for testing the result of a function.
	is := is.New(t)

	// is.NoErr will assert that the returned err is nil.
	// is.Equal will assert that the two values are equal.
	// is.True will assert that the comparison is true.
	// is.Fail will mark the test as failed.

	// For all of the asserts you can add a comment after that will be printed when the assert fails.

	r, err := fizzbuzz.FizzBuzz(1)
	is.NoErr(err)    // no error expected
	is.Equal(r, "1") // should be a number for non 3 or 5 divisible values.

	r, err = fizzbuzz.FizzBuzz(3)
	is.NoErr(err)
	is.Equal(r, "fizz") // should be fizz when divisible by 3.

	r, err = fizzbuzz.FizzBuzz(5)
	is.NoErr(err)
	is.Equal(r, "buzz") // should be buzz when divisible by 5.

	r, err = fizzbuzz.FizzBuzz(15)
	is.NoErr(err)
	is.Equal(r, "fizzbuzz") // should be fizzbuzz when divisible by 3 and 5.

	r, err = fizzbuzz.FizzBuzz(0)
	is.True(err != nil) // error expected for zerp.
	is.Equal(r, "")     // no value expected when error.

	// Testing panics can be tricky since we need to use a defer to recover and assert.
	// Panic should be a very rare occurance! This example is to demonstrate how it could be done.
	func() {
		defer func() {
			if p := recover(); p == nil {
				is.Fail() // a negative value should panic.
			}
		}()

		r, err := fizzbuzz.FizzBuzz(-1)
		// NOTE: These assertions will never get run because of the panic.
		is.NoErr(err)
		is.Equal(r, "")
	}()
}

// TestFizzBuzz_Table shows how the test can be extended to run multiple inputs with the same test code.
// This can be very useful as new test cases are added over time on a function.
func TestFizzBuzz_Table(t *testing.T) {
	// Create an anonymous struct
	// that has your expected in and out. To test error states include either a boolian to indicate it should
	// error or the error value to test for equality.
	tests := []struct {
		in      int
		out     string
		isErr   bool
		isPanic bool
	}{
		{in: 1, out: "1"},
		{in: 3, out: "fizz"},
		{in: 5, out: "buzz"},
		{in: 15, out: "fizzbuzz"},
		{in: -1, isPanic: true},
		{in: 0, isErr: true},

		// Examples that are designed to be caught by the test to prove that it is testing conditions.
		// {in: 27, out: "27"},
		// {in: 1, isErr: true },
		// {in: 1, isPanic: true},
		// {in: 0},
	}

	for i, tt := range tests {
		// in a for loop we sould use t.Run so that the sub-tests are printed as individual tests.
		// This helps in identifying which test has failed.
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			is := is.New(t)

			func() {
				defer func() {
					p := recover()
					switch {
					case tt.isPanic && p == nil:
						is.Fail() // call should have a panic

					case tt.isPanic && p != nil:
						return

					case !tt.isPanic && p == nil:
						return

					case !tt.isPanic && p != nil:
						is.Fail() // call should not have panic
					}
				}()

				fizzbuzz.FizzBuzz(tt.in)
			}()
			if tt.isPanic {
				return
			}

			r, err := fizzbuzz.FizzBuzz(tt.in)
			if tt.isErr {
				is.True(err != nil) // should have returned error
			} else {
				is.NoErr(err) // should not have returned error
			}
			is.Equal(r, tt.out) // should be equal
		})
	}
}

// TestFizzBuzz_Helper shows how a large test can be broken down into smaller sub-tests.
// This can be usefull to make the tests more clear and easy to follow. Also, the sub-tests
// can be reused by other tests.
func TestFizzBuzz_Helper(t *testing.T) {
	tests := []struct {
		story string // short description for the test.

		in int

		hasFizz bool
		hasBuzz bool

		isErr   bool
		isPanic bool
	}{
		{story: "non 3 or 5 divisor should return number", in: 1},
		{story: "3 divisor should return fizz", in: 3, hasFizz: true},
		{story: "5 divisor should return buzz", in: 5, hasBuzz: true},
		{story: "3 and 5 divisor should return fizzbuzz", in: 15, hasFizz: true, hasBuzz: true},
		{story: "zero should return error", in: 0, isErr: true},
		{story: "negative number should panic", in: -1, isPanic: true},

		// Examples that are designed to be caught by the test to prove that it is testing conditions.
		// {story: "FailTest 27 is divisible by 3", in: 27},
		// {story: "TestFail 1 returns error", in: 1, isErr: true},
		// {story: "TestFail 1 returns panic", in: 1, isPanic: true},
		// {story: "TestFail -1 returns no panic", in: -1},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d %s", i, tt.story), func(t *testing.T) {
			t.Parallel() // t.Parallel will cause the test to be split into a goroutine and run each test at the same time.

			is := is.New(t)

			testPanic(t, tt.in, tt.isPanic)
			if tt.isPanic {
				return // if testing for panic return so next test can be run.
			}
			testErr(t, tt.in, tt.isErr)
			if tt.isErr {
				return // if testing for error return so next test can be run.
			}

			if tt.hasFizz {
				is.True(testFizz(t, tt.in)) // should have fizz
			} else {
				is.True(!testFizz(t, tt.in)) // should not have fizz
			}

			if tt.hasBuzz {
				is.True(testBuzz(t, tt.in)) // should have buzz
			} else {
				is.True(!testBuzz(t, tt.in)) // should not have buzz
			}
		})
	}
}

func testFizz(t *testing.T, in int) bool {
	t.Helper() // helper informs test framework that this is a subtest.
	is := is.New(t)

	r, err := fizzbuzz.FizzBuzz(in)
	is.NoErr(err)
	return strings.HasPrefix(r, "fizz")
}

func testBuzz(t *testing.T, in int) bool {
	t.Helper()
	is := is.New(t)

	r, err := fizzbuzz.FizzBuzz(in)
	is.NoErr(err)
	return strings.HasSuffix(r, "buzz")
}

func testErr(t *testing.T, in int, isErr bool) {
	is := is.New(t)

	_, err := fizzbuzz.FizzBuzz(in)
	if isErr {
		is.True(err != nil) // should have returned an error
	} else {
		is.NoErr(err) // should not have returned an error
	}
}

func testPanic(t *testing.T, in int, isPanic bool) {
	is := is.New(t)

	defer func() {
		p := recover()
		switch {
		case isPanic && p == nil:
			is.Fail() // call should have a panic

		case isPanic && p != nil:
			return

		case !isPanic && p == nil:
			return

		case !isPanic && p != nil:
			is.Fail() // call should not have panic
		}
	}()

	fizzbuzz.FizzBuzz(in)
}
