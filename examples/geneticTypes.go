package main

import "fmt"

type Wrapper[T any] /* any or interface name */ struct {
	Value T
}

func main() {
	stringWrapper := Wrapper[string]{"string"}
	fmt.Println(stringWrapper.Value)
	intWrapper := Wrapper[int]{12}
	fmt.Println(intWrapper.Value)
}
