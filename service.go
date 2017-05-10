package stringsvc1

import (
	"errors"
	"strings"
)

// StringService - model service as an interface
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// **** Business logic

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// StringSrv - implements StringService
type StringSrv struct{}

// Uppercase - upper cases chars in string
func (StringSrv) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToUpper(s), nil
}

// Count - counts chars in string
func (StringSrv) Count(s string) int {
	return len(s)
}
