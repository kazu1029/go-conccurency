package main

import (
	"fmt"
	"sync"
)

func main() {
	dc := NewDoubleCounter()
	for i := 0; i < 1000; i++ {
		go func(t int) {
			dc.Add(t)
		}(i % 2)
	}

	dc.Wait()
	dc.Print()
}

type DoubleCounter struct {
	valueA int
	valueB int

	commandCh chan int
	wg        sync.WaitGroup
}

func NewDoubleCounter() *DoubleCounter {
	dc := &DoubleCounter{
		valueA:    0,
		valueB:    0,
		commandCh: make(chan int),
	}

	dc.wg.Add(1000)
	go dc.start()
	return dc
}

func (dc *DoubleCounter) start() {
	for {
		target := <-dc.commandCh
		if target == 0 {
			dc.valueA++
		} else {
			dc.valueB++
		}
		dc.wg.Done()
	}
}

func (dc *DoubleCounter) Add(t int) {
	dc.commandCh <- t
}

func (dc *DoubleCounter) Wait() {
	dc.wg.Wait()
}

func (dc *DoubleCounter) Print() {
	fmt.Printf("A: %d, B: %d\n", dc.valueA, dc.valueB)
}
