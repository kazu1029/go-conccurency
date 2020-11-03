package main

import (
	"context"
	"fmt"
	"time"
)

func infiniteLoop1(ctx context.Context) {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		fmt.Println("Help!")

		select {
		case <-innerCtx.Done():
			fmt.Println("Exit from hell.")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go infiniteLoop1(ctx)

	cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	time.Sleep(1 * time.Second)
}
