package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty int array:", a)

	a[4] = 100
	fmt.Println("Last element is", a[len(a)-1])

	// declard and initialize
	b := []int{1, 2, 3, 5, 8}
	fmt.Println("b is", b)

	twoD := [][]int{
		{1, 2, 3},
		{4, 5},
	}
	fmt.Println("here's a two dimensional array", twoD)
}
