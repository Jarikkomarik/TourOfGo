package main

import "fmt"

func main() { //alternative to lambdas
	myFun := func(a, b int) int { return a + b }
	fmt.Println(compute(myFun))

}

func compute(fun func(a, b int) int) int {
	return fun(3, 4)
}
