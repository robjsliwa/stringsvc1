package stringsvc1

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware - loggin middleware
type LoggingMiddleware struct {
	Logger log.Logger
	Next   StringService
}

// Uppercase - upper cases string
func (mw LoggingMiddleware) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Uppercase(ctx, s)
	return
}

// Count - counts number of chars in string
func (mw LoggingMiddleware) Count(ctx context.Context, s string) (n int, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took:", time.Since(begin),
		)
	}(time.Now())

	n, err = mw.Next.Count(ctx, s)
	return
}
