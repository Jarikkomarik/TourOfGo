package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChanel := make(chan int, 2) // init channel with size of 2 elements
	go produceVal(bufferedChanel)       // adds element each 2 seconds
	go produceVal(bufferedChanel)       // adds element each 2 seconds

	for {
		fmt.Println(<-bufferedChanel) //printing elements once present
	}

}

func produceVal(bufferedChannel chan int) {
	for {
		time.Sleep(2 * time.Second)
		bufferedChannel <- 1
	}
}
