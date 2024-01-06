package main

import "fmt"

type A struct {
	A int
}

func (a *A) incrementMethod(b int) {
	a.A = a.A + b
}

func (a A) PrintMethod() {
	fmt.Println(a.A)
}

func incrementFunc(a *A, b int) {
	a.A = a.A + b
}

func main() {
	a := A{10}
	fmt.Println(a.A)

	aPointer := &a
	aPointer.incrementMethod(10) // can call method frm line 9 on pointer
	fmt.Println(a.A)

	a.incrementMethod(10) // also can call same method on variable itself
	fmt.Println(a.A)

	//incrementFunc(a, 10) // won't compile because of incrementFunc expects pinter
	incrementFunc(aPointer, 10) // works as expected
	fmt.Println(a.A)
}
