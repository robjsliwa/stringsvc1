package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

// **** Requests and responses

type uppercaseRequest struct {
	InputString string `json:"input_string"`
}

type uppercaseResponse struct {
	UppercasedString string `json:"uppercased_string"`
	Err              string `json:"error_message,omitempty"`
}

type countRequest struct {
	InputString string `json:"input_string"`
}

type countResponse struct {
	Length int `json:"length"`
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	uppercaseHandler := httptransport.NewServer(
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
	http.Handle("/count", countHandler)
	logger.Log(http.ListenAndServe(":8080", nil))
}
