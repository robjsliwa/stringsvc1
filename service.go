package stringsvc1

import (
	"context"
	"errors"
	"strings"
)

// StringService - model service as an interface
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) (int, error)
}

// **** Business logic

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// StringSrv - implements StringService
type StringSrv struct{}

// Uppercase - upper cases chars in string
func (StringSrv) Uppercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToUpper(s), nil
}

// Count - counts chars in string
func (StringSrv) Count(ctx context.Context, s string) (int, error) {
	return len(s), nil
}
