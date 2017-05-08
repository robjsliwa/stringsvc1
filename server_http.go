package main

import (
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHTTPHandler - make http handler
func MakeHTTPHandler(svc StringService, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/uppercase").Handler(httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/count").Handler(httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	))

	return r
}
