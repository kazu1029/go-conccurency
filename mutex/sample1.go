package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(getNumber())
}

type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get() int {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

func (i *SafeNumber) Set(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func getNumber() int {
	i := &SafeNumber{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		i.Set(5)
	}()

	wg.Wait()
	return i.Get()
}
