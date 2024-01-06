package main

import "fmt"

func main() {
	_, b := ErrorIfOdd(1)
	fmt.Println(b) //printing error
}

type NotEvenError struct {
}

func (notEvenError NotEvenError) Error() string { //implementing error
	return "passed argument is not odd"
}

func ErrorIfOdd(a int) (int, error) {
	if a%2 == 0 {
		return a, nil
	} else {
		return a, NotEvenError{} //returning error
	}
}
