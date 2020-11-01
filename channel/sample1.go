package main

import (
	"fmt"
	"time"
)

func main() {
	// theMine := [5]string{"ore1", "ore2", "ore3"}
	// oreChan := make(chan string)

	// // Finder
	// go func(mine [5]string) {
	// 	for _, item := range mine {
	// 		oreChan <- item // send
	// 	}
	// }(theMine)

	// // Ore Breaker
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		foundOre := <-oreChan // receive
	// 		fmt.Println("Miner: Received " + foundOre + " from finder")
	// 	}
	// }()
	// <-time.After(5 * time.Second)

	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item
			}
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre"
		}
	}()

	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(5 * time.Second)
}
