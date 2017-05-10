package stringsvc1

import (
	"context"

	"github.com/robjsliwa/stringsvc1/pb"
)

// DecodeGRPCUppercaseRequest - decode uppercase request
func DecodeGRPCUppercaseRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UppercaseRequest)
	return pb.UppercaseRequest{InputString: req.InputString}, nil
}

// EncodeGRPCUppercaseResponse - encode uppercase response
func EncodeGRPCUppercaseResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.UppercaseResponse)
	return &pb.UppercaseResponse{UppercasedString: res.UppercasedString, Err: res.Err}, nil
}

// DecodeGRPCCountRequest - decode count request
func DecodeGRPCCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CountRequest)
	return pb.CountRequest{InputString: req.InputString}, nil
}

// EncodeGRPCCountResponse - encode count response
func EncodeGRPCCountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CountResponse)
	return &pb.CountResponse{Length: res.Length, Err: res.Err}, nil
}

// --- Client

// EncodeGRPCUppercaseRequest - encode uppercase request
func EncodeGRPCUppercaseRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UppercaseRequest)
	return &pb.UppercaseRequest{InputString: req.InputString}, nil
}

// DecodeGRPCUppercaseResponse - decode uppercase response
func DecodeGRPCUppercaseResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.UppercaseResponse)
	return pb.UppercaseResponse{UppercasedString: res.UppercasedString, Err: res.Err}, nil
}

// EncodeGRPCCountRequest - encode count request
func EncodeGRPCCountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CountRequest)
	return &pb.CountRequest{InputString: req.InputString}, nil
}

// DecodeGRPCCountResponse - decode count response
func DecodeGRPCCountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CountResponse)
	return pb.CountResponse{Length: res.Length, Err: res.Err}, nil
}
