package main

import "fmt"

func main() {

	//buf := "test value"
	buf := "test \t value" + " привет!"

	//num := 5
	//var num1 int

	//num, num1 := 35, "thirty five"

	//num, num1 := 1.001, 2.333
	//var num bool = true
	//num1 := true

	//const (
	//	num  = "hello"
	//	num1 = 35
	//)
	const (
		num = iota
		_
		_
		num1
	)

	fmt.Println(buf[:3], num, num1)

}
