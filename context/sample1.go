package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	ctx = context.WithValue(ctx, "message", "hi")
	defer cancel()

	go infiniteLoop(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func infiniteLoop(ctx context.Context) {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		fmt.Println("Waiting for time out!")
		fmt.Println("message:", ctx.Value("message").(string))
		select {
		case <-innerCtx.Done():
			fmt.Println("Exit now!")
			return
		default:
			fmt.Println("default")
		}
	}
}
