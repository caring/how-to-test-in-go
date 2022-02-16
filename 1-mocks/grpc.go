package mocks

import (
	"context"

	errors "github.com/caring/gopkg-errors"
	"google.golang.org/grpc/codes"

	fizzbuzz "github.com/caring/test/0-begin-here"
	"github.com/caring/test/1-mocks/pb"
)

type GRPCService struct {
	pb.UnimplementedMockServiceServer
}

func (GRPCService) GetFizzBuzz(ctx context.Context, in *pb.GetFizzBuzzRequest) (*pb.GetFizzBuzzResponse, error) {
	fb, err := fizzbuzz.FizzBuzz(int(in.GetNumber()))
	if err != nil {
		return nil, errors.WithGrpcStatus(err, codes.Internal)
	}

	return &pb.GetFizzBuzzResponse{
		FizzBuzz: fb,
	}, nil
}

func GetFizzBuzzGRPC(ctx context.Context, n uint64, cl pb.MockServiceClient) (string, error) {
	res, err := cl.GetFizzBuzz(ctx, &pb.GetFizzBuzzRequest{Number: n})
	if err != nil {
		return "", err
	}
	return res.GetFizzBuzz(), nil
}
