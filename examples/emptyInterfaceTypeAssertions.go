package main

import "fmt"

type HelloAble interface {
	DoHello()
}

type HelloAbleImpl struct {
}

func (HelloAbleImpl HelloAbleImpl) DoHello() {
	println("hello")
}

func main() {
	var helloAbleImpl interface{} = HelloAbleImpl{} //hides the type
	i, ok := helloAbleImpl.(HelloAbleImpl)          // casting to type in ()
	//i := helloAbleImpl.(HelloAbleImpl)// alternative, returns panic in case of unsuccessful casting

	fmt.Println(i)
	fmt.Println(ok)

	x, y := helloAbleImpl.(string) //unsuccessful casting example

	fmt.Println(x)
	fmt.Println(y)

	//type assertion switch example

	switch v := helloAbleImpl.(type) {
	case HelloAble:
		fmt.Println("helloAble", v)
	case int:
		fmt.Println("int", v)
	default:
		fmt.Println("default", v)
	}
	// todo: how to access v after switch?

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}
