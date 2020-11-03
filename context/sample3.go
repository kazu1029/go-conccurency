package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	go parentProcess(ctx)
	time.Sleep(20 * time.Second)
}

func parentProcess(ctx context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	childCtx2, _ := context.WithCancel(childCtx)
	go childProcess(childCtx, "child1")
	go childProcess(childCtx2, "child2")
	time.Sleep(10 * time.Second)
	cancel()
	ctxWithTimeout10, _ := context.WithTimeout(ctx, time.Second*10)
	go childProcess(ctxWithTimeout10, "with timeout")
}

func childProcess(ctx context.Context, prefix string) {
	for i := 1; i <= 1000; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: canceled \n", prefix)
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("%s:%d sec..\n", prefix, i)
		}
	}
}
