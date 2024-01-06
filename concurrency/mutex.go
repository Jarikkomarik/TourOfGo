package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentWrapper struct {
	mutex sync.Mutex // sync lock that can be used on object
	value int
}

func main() {

	wrapper := ConcurrentWrapper{sync.Mutex{}, 0}

	go increment10Times(&wrapper)
	go increment10Times(&wrapper)

	time.Sleep(time.Second)

}

func (wrapper *ConcurrentWrapper) increment() {
	wrapper.mutex.Lock() //blocks thread
	wrapper.value = wrapper.value + 1
	fmt.Println(wrapper.value)
	wrapper.mutex.Unlock() //unblocks
}

func increment10Times(wrapper *ConcurrentWrapper) {
	for i := 0; i < 10; i++ {
		wrapper.increment()
	}

}
