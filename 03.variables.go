package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// type inference
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println("zero value for an int is", e)

	// declare and initialize shorthand
	f := "apple"
	// var f string = "apple"
	fmt.Println("F for", f)
}
