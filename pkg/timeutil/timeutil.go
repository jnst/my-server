package timeutil

import (
	"context"
	"time"
)

type contextKey = struct{}

var timeKey = contextKey{}

func NowWithContext(ctx context.Context, now time.Time) context.Context {
	return context.WithValue(ctx, timeKey, now)
}

func Now(ctx context.Context) time.Time {
	t, ok := ctx.Value(timeKey).(time.Time)
	if !ok {
		return time.Now()
	}
	return t
}
