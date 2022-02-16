/* Mock HTTP


Additional Reaing
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
	// Setup Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/fizzbuzz/") {
			t.Errorf("expected to request '/fizzbuzz/', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"fizz_buzz":"fizzbuzz"}`))
	}))
	defer server.Close()

	// run client request
	value, _ := mocks.GetFizzBuzzHTTP(context.TODO(), server.URL, 15)

	// validate output
	if value != "fizzbuzz" {
		t.Errorf("expected 'fixed', got %s", value)
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
