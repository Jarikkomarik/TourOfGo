package main

import "fmt"

func main() {
	var jarik = User{Age: 12, Name: "Jarik", Surname: "Komarik", IsRegistered: false}
	jarik.IsRegistered = true

	var boris = User{"Boris", "Keke", 12, true}
	borisPointer := &boris
	borisPointer.Age = -1

	borisRef := boris
	borisRef.Name = "NotBoris"

	var userArr = []User{jarik, boris, borisRef}

	fmt.Println(userArr)
}

type User struct {
	Name         string
	Surname      string
	Age          int
	IsRegistered bool
}
