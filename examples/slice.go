package main

import "fmt"

func main() {

	var array = [4]int{1, 2, 3, 4} //array of fixed length

	fmt.Println(array)

	var slice = array[0:2] //from : to (exclusive) * just a references to elements from array
	slice = array[:2]      //from beginning : to (exclusive)

	fmt.Println(slice)

	array[0] = -1 //updated value of array

	fmt.Println(slice) //slice also updated

	//////////////////

	var sliceDynSize = []int{} //not declaring size (in this case it creates an empty array)
	for i := 0; i < 10; i++ {
		sliceDynSize = append(sliceDynSize, i)                                        // creates new array sliceDynSize + new element
		fmt.Printf("length %d, capacity %d \n", len(sliceDynSize), cap(sliceDynSize)) //size grows when new element is added, capacity is multiplied when limit is reached (1, 2, 4, 16, 32....)
	}

	fmt.Println(sliceDynSize)

	////////////////         type, size, capacity
	var sliceWithMake = make([]int, 10, 11)

	fmt.Printf("length %d, capacity %d \n", len(sliceWithMake), cap(sliceWithMake))

	//////////////// nested slices

	var nestedSlice = [][]string{{"X", "O", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}

	for i := 0; i < len(nestedSlice); i++ {
		for y := 0; y < len(nestedSlice[i]); y++ {
			fmt.Print(nestedSlice[i][y])
		}
		fmt.Println()
	}

	///////////// iterate over slice using "range"

	var rangeExampleSlice = []int{1, 2, 3, 4}

	for x, j := range rangeExampleSlice {
		fmt.Printf("index %d, value %d. \n", x, j)
	}

	// it is possible to use range with one var only. in this case it will be the index.

	for j := range rangeExampleSlice {
		fmt.Printf("value %d. \n", j)
	}

	//also following format is accepted
	//for i, _ := range pow
	//for _, value := range pow
}
