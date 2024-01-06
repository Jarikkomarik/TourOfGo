package main

import (
	"fmt"
	"time"
)

func main() {
	go printString("longSleep", 100) // starts new thread
	printString("shortSleep", 10)
	time.Sleep(time.Duration(100000000))
	fmt.Println(123)
}

func printString(s string, sleepTime int) {
	time.Sleep(time.Duration(sleepTime))
	fmt.Println(s)
}
