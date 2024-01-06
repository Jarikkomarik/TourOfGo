package main

import "fmt"

type Printable interface { // declaring interface
	print()
}

type Kek struct {
	Name string
}

type Cheburek struct {
	Name string
}

func (kek *Kek) print() { //implementing interface by declaring methods of Printable
	if kek == nil {

	} else {
		println("Type (kek): " + kek.Name)
	}

}

func (kek Cheburek) print() {
	println("Type (Cheburek): " + kek.Name)
}

func main() {
	var a Printable
	a = &Kek{"kek"}
	a.print()

	a = Cheburek{"cheburek"}
	a.print()

	var kek = &Kek{"sdf"}
	describe(kek)

	var c Printable //operation with nil
	var g Kek
	c = &g
	c.print()
}

func describe(i Printable) {
	fmt.Printf("(%v, %T)\n", i, i)
}
