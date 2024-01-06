package main

import "fmt"

type Pair struct {
	A int
	B int
}

func (p Pair) printPair() { //by declaring (p Pair) we specify that it will belong to Pair struct * can be declared only in the same package as a struct
	fmt.Printf("%d : %d\n", p.A, p.B)
}

func printPair(pair Pair) { // method of this file that takes Pair as argument
	fmt.Printf("%d : %d\n", pair.A, pair.B)
}

////// following example shows difference between pointer arguments and value arguments.

func incrementValue(pair Pair) {
	pair.A = pair.A + 1
	pair.B = pair.B + 1
}

func incrementPointer(pair *Pair) {
	pair.A = pair.A + 1
	pair.B = pair.B + 1
}

func main() {
	pair := Pair{10, 10}
	incrementValue(pair) //actual pair is not changed because operation was done on copy of "pair"
	printPair(pair)
	incrementPointer(&pair) //pair values were changed trough pointer
	printPair(pair)
}
