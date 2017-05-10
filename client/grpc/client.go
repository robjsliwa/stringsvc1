package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/robjsliwa/stringsvc1"
	"github.com/robjsliwa/stringsvc1/pb"
	"google.golang.org/grpc"
)

// New - create new client
func New(conn *grpc.ClientConn) pb.Service {
	var uppercaseEndpoint = grpctransport.NewClient(
		conn, "StringService", "Uppercase",
		stringsvc1.EncodeGRPCUppercaseRequest,
		stringsvc1.DecodeGRPCUppercaseResponse,
		pb.UppercaseResponse{},
	).Endpoint()
	var countEndpoint = grpctransport.NewClient(
		conn, "StringService", "Count",
		stringsvc1.EncodeGRPCCountRequest,
		stringsvc1.DecodeGRPCCountResponse,
		pb.CountResponse{},
	).Endpoint()
	return vault.Endpoints{
		UppercaseEndpoint: uppercaseEndpoint,
		CountEndpoint:     countEndpoint,
	}
}
