package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000
	const d = 2e20 / n
	fmt.Println(d)

	fmt.Println("d explicitly converted to 64bit integer", int64(d))

	fmt.Println(math.Sin(n))
}
