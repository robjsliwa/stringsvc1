package grpcclient

import (
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/robjsliwa/stringsvc1"
	"github.com/robjsliwa/stringsvc1/pb"
	"google.golang.org/grpc"
)

// New - create new client
func New(conn *grpc.ClientConn) stringsvc1.StringService {
	var uppercaseEndpoint endpoint.Endpoint
	{
		uppercaseEndpoint = grpctransport.NewClient(
			conn, "StringService", "Uppercase",
			stringsvc1.EncodeGRPCUppercaseRequest,
			stringsvc1.DecodeGRPCUppercaseResponse,
			pb.UppercaseResponse{},
		).Endpoint()
	}

	var countEndpoint endpoint.Endpoint
	{
		countEndpoint = grpctransport.NewClient(
			conn, "StringService", "Count",
			stringsvc1.EncodeGRPCCountRequest,
			stringsvc1.DecodeGRPCCountResponse,
			pb.CountResponse{},
		).Endpoint()
	}

	return stringsvc1.Endpoints{
		Uppercase: uppercaseEndpoint,
		Count:     countEndpoint,
	}
}
