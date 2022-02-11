package mocks_test

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"

	mocks "github.com/caring/test/1-mocks"
	"github.com/caring/test/1-mocks/pb"
)

// TestFizzBuzzHandler shows how to test the server handler with some simulated GRPC requests.
func TestFizzBuzzGRPCHandler(t *testing.T) {
	tests := []struct {
		Number   uint64
		Code     codes.Code
		FizzBuzz string
		Error    string
	}{
		{1, codes.OK, "1", ""},
		{3, codes.OK, "fizz", ""},
		{5, codes.OK, "buzz", ""},
		{15, codes.OK, "fizzbuzz", ""},
	}

	for i, tt := range tests {
		t.Logf("Test %d - %v", i, tt.Number)

		c := &mocks.GRPCService{}

		result, err := c.GetFizzBuzz(context.TODO(), &pb.FizzBuzzRequest{Number: tt.Number})
		if err != nil {
			t.Errorf(err.Error())
		}

		if tt.FizzBuzz != result.GetFizzBuzz() {
			t.Errorf("expected fizzbuzz '%s' got '%s'", tt.FizzBuzz, result.GetFizzBuzz())
		}
	}
}


