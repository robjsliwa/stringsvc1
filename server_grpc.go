package stringsvc1

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/robjsliwa/stringsvc1/pb"
)

type grpcServer struct {
	uppercase grpctransport.Handler
	count     grpctransport.Handler
}

func (s *grpcServer) Uppercase(ctx context.Context, r *pb.UppercaseRequest) (*pb.UppercaseResponse, error) {
	_, resp, err := s.uppercase.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UppercaseResponse), nil
}

func (s *grpcServer) Count(ctx context.Context, r *pb.CountRequest) (*pb.CountResponse, error) {
	_, resp, err := s.count.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CountResponse), nil
}

// MakeGRPCServer - make gRPC server
func MakeGRPCServer(svc StringService, logger log.Logger) pb.StringServiceServer {
	return &grpcServer{
		uppercase: grpctransport.NewServer(
			makeUppercaseEndpoint(svc),
			DecodeGRPCUppercaseRequest,
			EncodeGRPCUppercaseResponse,
		),
		count: grpctransport.NewServer(
			makeCountEndpoint(svc),
			DecodeGRPCCountRequest,
			EncodeGRPCCountResponse,
		),
	}
}
