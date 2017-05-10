package stringsvc1

import (
	"context"
	"errors"

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
		uppercasedString, err := svc.Uppercase(ctx, req.InputString)
		if err != nil {
			return uppercaseResponse{uppercasedString, err.Error()}, nil
		}

		return uppercaseResponse{uppercasedString, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		length, _ := svc.Count(ctx, req.InputString)
		return countResponse{length}, nil
	}
}

// Endpoints - endpoints for primarily client
type Endpoints struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndpoint     endpoint.Endpoint
}

// Uppercase - upper case chars in string
func (e Endpoints) Uppercase(ctx context.Context, inputString string) (string, error) {
	req := uppercaseRequest{InputString: inputString}
	resp, err := e.UppercaseEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	uppercaseResp := resp.(uppercaseResponse)
	if uppercaseResp.Err != "" {
		return "", errors.New(uppercaseResp.Err)
	}
	return uppercaseResp.UppercasedString, nil
}

// Count - count chars in input string
func (e Endpoints) Count(ctx context.Context, inputString string) (int, error) {
	req := countRequest{InputString: inputString}
	resp, err := e.CountEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	countResp := resp.(countResponse)

	return countResp.Length, nil
}
