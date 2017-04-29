package main

import (
	"log"
	"net/http"

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
	svc := stringService{}

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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
