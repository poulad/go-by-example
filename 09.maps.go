package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"] = 1
	m["bar"] = 22

	fmt.Println(m)

	var v1 int
	v1 = m["foo"]
	fmt.Println("v1:", v1, "out of", len(m), "elements")

	delete(m, "foo")
	fmt.Println(m)

	key := "qux"
	_, isPresent := m[key]
	fmt.Println("key", key, "is present?", isPresent)

	// declare and init at once
	myMap := map[int]bool{0: false, 1: true}
	fmt.Println("myMap:", myMap)
}
