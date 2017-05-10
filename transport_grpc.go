package main

import (
	"context"

	"github.com/robjsliwa/stringsvc1/pb"
)

func decodeGRPCUppercaseRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UppercaseRequest)
	return pb.UppercaseRequest{InputString: req.InputString}, nil
}

func encodeGRPCUppercaseResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.UppercaseResponse)
	return &pb.UppercaseResponse{UppercasedString: res.UppercasedString, Err: res.Err}, nil
}

func decodeGRPCCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CountRequest)
	return pb.CountRequest{InputString: req.InputString}, nil
}

func encodeGRPCCountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CountResponse)
	return &pb.CountResponse{Length: res.Length, Err: res.Err}, nil
}
