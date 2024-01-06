package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan time.Time) // creating channel
	go getCurrentTime(c)      //passing channel to a function
	createdTime := <-c        //blocking main() until c channel val is not created
	fmt.Println(createdTime)  //printing produced var
}

func getCurrentTime(c chan time.Time) {
	time.Sleep(time.Duration(4) * time.Second) // sleeps for 4 sec
	c <- time.Now()                            //passing value to channel
}
