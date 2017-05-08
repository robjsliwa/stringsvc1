package main

import (
	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/examples/addsvc/pb"
)

// MakeGRPCServer - make gRPC server
func MakeGRPCServer(endpoints Endpoints, tracer stdopentracing.Tracer, logger log.Logger) pb.AddServer {
	return &grpcServer{
		uppercase: grpctransport.NewServer(
			makeUppercaseEndpoint(svc),
			decodeUppercaseRequest,
			encodeResponse,
		),
		count: grpctransport.NewServer(
			makeCountEndpoint(svc),
			decodeCountRequest,
			encodeResponse,
		),
	}
}

type grpcServer struct {
	uppercase grpctransport.Handler
	count     grpctransport.Handler
}
