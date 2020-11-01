package main

import "fmt"

func main() {
	fmt.Println(getNumber())

	i := <-getNumberChan()
	fmt.Println(i)
}

func getNumber() int {
	var i int
	done := make(chan struct{})
	go func() {
		i = 5
		done <- struct{}{}
	}()

	<-done
	return i
}

func getNumberChan() <-chan int {
	c := make(chan int)
	go func() {
		c <- 5
	}()
	return c
}
