package main

import (
	"fmt"
	"time"
)

func main() {
	valChan := make(chan string)
	quitChan := make(chan bool)

	go produce(valChan, quitChan)

	for i := 0; i < 10; i++ {
		fmt.Println(<-valChan)
		time.Sleep(1 * time.Second)
	}
	quitChan <- true

}

func produce(val chan string, quit chan bool) {
	for {
		select { // alternative of switch (waiting for first present value in channel)
		case val <- "x":
			fmt.Println("produced new val")
		case <-quit:
			fmt.Println("finished produce")
			break
		default:
			fmt.Println("sleep ....")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
