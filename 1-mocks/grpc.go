package mocks

import (
	"context"

	errors "github.com/caring/gopkg-errors"
	"google.golang.org/grpc/codes"

	fizzbuzz "github.com/caring/test/0-begin-here"
	"github.com/caring/test/1-mocks/pb"
)

type GRPCService struct{}

func (GRPCService) GetFizzBuzz(ctx context.Context, in *pb.FizzBuzzRequest) (*pb.FizzBuzzResponse, error) {
	fb, err := fizzbuzz.FizzBuzz(int(in.GetNumber()))
	if err != nil {
		return nil, errors.WithGrpcStatus(err, codes.Internal)
	}

	return &pb.FizzBuzzResponse{
		FizzBuzz: fb,
	}, nil
}
