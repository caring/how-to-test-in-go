/* Mock HTTP


Additional Reading
[1]: https://medium.com/zus-health/mocking-outbound-http-requests-in-go-youre-probably-doing-it-wrong-60373a38d2aa
*/
package mocks_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "github.com/caring/test/1-mocks"
)

// TestGetFizzBuzz shows how to test a client that calls an HTTP server.
// httptest.NewServer will modify the global RoundTripper until server.Close() is called.
func TestGetFizzBuzz(t *testing.T) {
	tests := []struct {
		// values sent to server
		URL string
		// values returned by server
		Code int
		Body []byte
		// expected values
		Expected string
		Error    string
	}{
		{"/fizzbuzz/1", http.StatusOK, []byte(`{"fizz_buzz":"1"}`), "1", ""},
		{"/fizzbuzz/3", http.StatusOK, []byte(`{"fizz_buzz":"fizz"}`), "fizz", ""},
		{"/fizzbuzz/5", http.StatusOK, []byte(`{"fizz_buzz":"buzz"}`), "buzz", ""},
		{"/fizzbuzz/13", http.StatusOK, []byte(`{"fizz_buzz":"fizzbuzz"}`), "fizzbuzz", ""},
		{"/fizzbuzz/0", http.StatusInternalServerError, []byte(`{"fizz_buzz":""}`), "", "failed to read response. got 500"},
		{"/fizzbuzz/1", http.StatusOK, []byte(`{"fizz_buzz":"`), "", "unexpected end of JSON input"},
	}

	for _, tt := range tests {
		// Setup Mock HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/fizzbuzz/") {
				t.Errorf("expected to request '/fizzbuzz/', got: %s", r.URL.Path)
			}
			if r.Header.Get("Accept") != "application/json" {
				t.Errorf("expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
			}
			w.WriteHeader(tt.Code)
			w.Write(tt.Body)
		}))
		defer server.Close()

		// run client request
		value, err := mocks.GetFizzBuzzHTTP(context.TODO(), server.URL, 15)

		if tt.Error == "" && err != nil {
			t.Errorf(err.Error())
		}

		if tt.Error != "" && err != nil && tt.Error != err.Error() {
			t.Errorf("exptected %q, got %q", tt.Error, err.Error())
		}

		// validate output
		if value != tt.Expected {
			t.Errorf("expected %q, got %q", tt.Expected, value)
		}
	}
}

// TestFizzBuzzHandler shows how to test the server handler with some simulated HTTP requests.
func TestFizzBuzzHandler(t *testing.T) {
	// setup a table test for requests.
	tests := []struct {
		Method   string
		URL      string
		Accept   string
		Code     int
		FizzBuzz string
		Error    string
	}{
		{http.MethodPost, "/", "text/plain", http.StatusMethodNotAllowed, "", "not allowed: POST"},
		{http.MethodGet, "/", "text/plain", http.StatusNotFound, "", "not found: /"},
		{http.MethodGet, "/fizzbuzz/1", "text/plain", http.StatusNotAcceptable, "", "not acceptable: text/plain"},
		{http.MethodGet, "/fizzbuzz/abc", "application/json", http.StatusBadRequest, "", "bad request: abc"},
		{http.MethodGet, "/fizzbuzz/1", "application/json", http.StatusOK, "1", ""},
		{http.MethodGet, "/fizzbuzz/3", "application/json", http.StatusOK, "fizz", ""},
		{http.MethodGet, "/fizzbuzz/5", "application/json", http.StatusOK, "buzz", ""},
		{http.MethodGet, "/fizzbuzz/15", "application/json", http.StatusOK, "fizzbuzz", ""},
		{http.MethodGet, "/fizzbuzz/0", "application/json", http.StatusInternalServerError, "", "internal error: too much fizzbuzzery"},
	}

	for i, tt := range tests {
		t.Logf("Test %d - %s %s", i, tt.Method, tt.URL)

		// build the mocked request and recorder
		r := httptest.NewRequest(tt.Method, tt.URL, nil)
		r.Header.Set("Accept", tt.Accept)
		w := httptest.NewRecorder()

		// call handler
		mocks.FizzBuzzHandler(w, r)

		// check the http response code is correct.
		if tt.Code != w.Code {
			t.Errorf("expected '%d' got '%d'", tt.Code, w.Code)
		}

		// decode the json payload
		var result struct {
			FizzBuzz string `json:"fizz_buzz"`
			Error    string `json:"err"`
		}
		json.NewDecoder(w.Body).Decode(&result)

		// check error message
		if tt.Error != result.Error {
			t.Errorf("expected error '%s' got '%s'", tt.Error, result.Error)
		}
		// check returned value
		if tt.FizzBuzz != result.FizzBuzz {
			t.Errorf("expected fizzbuzz '%s' got '%s'", tt.FizzBuzz, result.FizzBuzz)
		}
	}
}
