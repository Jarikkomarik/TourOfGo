package main

import "fmt"

var a, kek, lol = true, 0.12, "sdfg"

func main() {

	i := 10

	x := &i

	println(*x)
	i = 1
	println(*x)
	*x = 12

	println(i)

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

}

func sum(a, b int) int {
	return a + b
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "hey, u r negative bruh ://///"
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		var z = float64(10)
		var temp = z
		for {
			z -= (z*z - x) / (2 * z)
			if temp == z {
				break
			} else {
				temp = z
			}

		}
		return z, nil
	}

}
