package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
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

// **** Endpoints

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		uppercasedString, err := svc.Uppercase(req.InputString)
		if err != nil {
			return uppercaseResponse{uppercasedString, err.Error()}, nil
		}

		return uppercaseResponse{uppercasedString, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		length := svc.Count(req.InputString)
		return countResponse{length}, nil
	}
}
