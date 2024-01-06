package main

import "fmt"

func main() {

	aMap := map[string]int{
		"jarik":   1,
		"komarik": 2,
	}
	fmt.Println(aMap)

	myMap := make(map[string]int)
	myMap["jarik"] = 12
	fmt.Println(myMap)

	structMap := map[string]UserKek{
		"jarik": {12, false},
		"pepa":  {2, true},
	}

	fmt.Println(structMap)

	var value = structMap["jarik"] // get value
	fmt.Println(value)
	value, isPresent := structMap["shit"] // check whether is present
	fmt.Println(isPresent)

	delete(structMap, "pepa") // delete element
	fmt.Println(structMap)
}

type UserKek struct {
	age          int
	isRegistered bool
}
