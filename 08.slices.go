// go run ./08.slices.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := []int{0, 0, 0}
	slc := make([]int, 3)
	// what's the diff????!!!
	fmt.Println("array:", reflect.TypeOf(arr), arr)
	fmt.Println("slice:", reflect.TypeOf(slc), slc)

	s := make([]string, 3)
	fmt.Println("empty slice is", s)
	fmt.Printf("s is of type %T and size %d\n", reflect.TypeOf(s).Name(), len(s))

	slc = append(slc, 444, 55, 6)
	fmt.Println("our slice is now", slc)

	copiedSlc := make([]int, len(slc)+3)
	numOfElements := copy(copiedSlc, slc)
	fmt.Println("Copied %d elements from the original slice", numOfElements)

	fmt.Println("Here's an slice, positions 2 to 6, of the cloned slice:", copiedSlc[2:7])

	twoDimSlice := [][]string{{"a", "b", "c"}, {}, {"foo", "bar"}, {"x"}}
	fmt.Println("2-dimensional slice:", twoDimSlice)
}
