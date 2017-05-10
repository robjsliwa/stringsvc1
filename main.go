package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/robjsliwa/stringsvc1/pb"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	/*uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)*/

	errc := make(chan error)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var h http.Handler
	{
		h = MakeHTTPHandler(svc, logger)
	}

	// HTTP transport.
	go func() {
		logger = log.With(logger, "transport", "HTTP")
		//logger.Log("addr", *httpAddr)
		errc <- http.ListenAndServe(":8080", h)
	}()

	// gRPC transport.
	go func() {
		logger = log.With(logger, "transport", "gRPC")

		//ln, err := net.Listen("tcp", *grpcAddr)
		ln, err := net.Listen("tcp", ":8081")
		if err != nil {
			errc <- err
			return
		}

		srv := MakeGRPCServer(svc, logger)
		s := grpc.NewServer()
		pb.RegisterStringServiceServer(s, srv)

		//logger.Log("addr", *grpcAddr)
		errc <- s.Serve(ln)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
