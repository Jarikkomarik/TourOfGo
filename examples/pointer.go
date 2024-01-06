package main

func main() {
	i := 10

	x := &i //assign pointer of i to x

	println(*x) // print value that x refers to
	i = 1       // change value of i
	println(*x) // print value that x refers to
	*x = 12     // change value of i through pointer x

	println(i)
}
