package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	fizzbuzz "github.com/caring/test/0-begin-here"
)

// GetFizzBuzzHTTP makes an HTTP request to a fizzbuzz service an returns the result.
func GetFizzBuzzHTTP(ctx context.Context, baseURL string, n int) (string, error) {
	url := fmt.Sprintf("%s/fizzbuzz/%d", baseURL, n)

	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	request.Header.Add("Accept", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("failed to read response. got %d", response.StatusCode)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	v := &struct {
		FizzBuzz string `json:"fizz_buzz"`
	}{}

	err = json.Unmarshal(content, v)
	if err != nil {
		return "", err
	}

	return v.FizzBuzz, nil
}

// FizzBuzzHandler is a simple HTTP handler that returns a fizzbuzz result.
// The encoding errors should probably be logged, but were left out of this example.
func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	var result struct {
		FizzBuzz string `json:"fizz_buzz"`
		Error    string `json:"err"`
	}
	enc := json.NewEncoder(w)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		result.Error = "not allowed: " + r.Method
		_ = enc.Encode(result)

		return
	}

	if !strings.HasPrefix(r.URL.Path, "/fizzbuzz/") {
		w.WriteHeader(http.StatusNotFound)
		result.Error = "not found: " + r.URL.Path
		_ = enc.Encode(result)

		return
	}

	if r.Header.Get("Accept") != "application/json" {
		w.WriteHeader(http.StatusNotAcceptable)
		result.Error = "not acceptable: " + r.Header.Get("Accept")
		_ = enc.Encode(result)

		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/fizzbuzz/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result.Error = "bad request: " + idStr
		_ = enc.Encode(result)

		return
	}

	result.FizzBuzz, err = fizzbuzz.FizzBuzz(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		result.Error = "internal error: " + err.Error()
		_ = enc.Encode(result)

		return
	}

	w.WriteHeader(http.StatusOK)
	_ = enc.Encode(result)
}
