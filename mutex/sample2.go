package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	wg := &sync.WaitGroup{}
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(3 * time.Second)
	fmt.Println(counter)
}
