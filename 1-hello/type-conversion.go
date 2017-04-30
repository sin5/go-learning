package main

import (
	"math"
	"fmt"
)

func main() {
	var x, y int = 3, 4
	f := math.Sqrt(float64(x * x + y * y))
	z := uint(f)

	fmt.Println(x, y, f, z)
}