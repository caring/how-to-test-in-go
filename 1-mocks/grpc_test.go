package mocks_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	logging "github.com/caring/gopkg-logging"
	muxserver "github.com/caring/gopkg-muxserver"

	mocks "github.com/caring/test/1-mocks"
	"github.com/caring/test/1-mocks/pb"
)

// TestFizzBuzzHandler shows how to test the server handler with some simulated GRPC requests.
func TestFizzBuzzGRPCHandler(t *testing.T) {
	// table test for requests
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
		{0, codes.Internal, "", "rpc error: code = Internal desc = too much fizzbuzzery"},
	}

	for i, tt := range tests {
		t.Logf("Test %d - %v", i, tt.Number)

		// instantiate the grpc service
		c := &mocks.GRPCService{}

		// Execute handler
		result, err := c.GetFizzBuzz(context.TODO(), &pb.GetFizzBuzzRequest{Number: tt.Number})

		if tt.Error == "" && err != nil {
			t.Errorf(err.Error())
		}
		if tt.Error != "" && tt.Error != err.Error() {
			t.Errorf("expected: %v, got: %v", tt.Error, err.Error())
		}


		// Test result
		if tt.FizzBuzz != result.GetFizzBuzz() {
			t.Errorf("expected fizzbuzz '%s' got '%s'", tt.FizzBuzz, result.GetFizzBuzz())
		}
	}
}

// TestFizzBuzzGRPCClient simulates actual calls to a mock GRPC server to test the client.
func TestFizzBuzzGRPCClient(t *testing.T) {
	// table test for requests
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
		{0, codes.Internal, "", "rpc error: code = Internal desc = too much fizzbuzzery"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d - %v", i, tt.Number), func(t *testing.T) {
			// Setup context for request.
			ctx, stop := context.WithCancel(context.Background())
			defer stop() // teardown server at end of test.

			// Setup GRPC server
			cl, _ := getGRPCServerClient(ctx, t)

			n, err := mocks.GetFizzBuzzGRPC(ctx, tt.Number, cl)

			if tt.Error != "" && err != nil && tt.Error != err.Error() {
				t.Errorf("expected error '%s' got '%s'", tt.Error, err.Error())
			}

			if n != tt.FizzBuzz {
				t.Errorf("expected fizzbuzz '%s' got '%s'", tt.FizzBuzz, n)
			}

		})
	}
}

// getGRPCServerClient starts a server using a bufcon connection and returns the client.
func getGRPCServerClient(ctx context.Context, t *testing.T) (pb.MockServiceClient, func()) {
	t.Helper()

	// make sure we dont hang forever in CI
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	lis := bufconn.Listen(1024 * 1024)
	srv := muxserver.New(logging.NewNopLogger(), nil, lis)

	// Add GRPC service handler
	pb.RegisterMockServiceServer(srv.GRPC, &mocks.GRPCService{})

	// Start service
	go func() {
		defer lis.Close()
		err := srv.Serve()
		if err != nil {
			t.Log(err)
		}
	}()

	// Test GRPC connection
	conn, _ := grpc.DialContext(
		ctx, "bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) { return lis.Dial() }),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	return pb.NewMockServiceClient(conn), cancel
}
