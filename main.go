package main

import (
	"context"
	"fmt"
	"my-server/pkg/timeutil"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = timeutil.NowWithContext(ctx, time.Now())

	fmt.Println(timeutil.Now(ctx).Format(time.RFC3339))
}
