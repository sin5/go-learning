package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1
	i complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, i, i)
}