package main

import "fmt"

// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
func main() {
	c := make(chan int)
	go produce100values(c)
	for {
		val, isWorking := <-c
		if isWorking {
			fmt.Println(val)
		} else { //when closed
			break
		}
	}
	fmt.Println("finish")
}

func produce100values(c chan int) { //add 100 records and close chan
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // closing channel
}
