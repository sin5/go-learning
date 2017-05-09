package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	m := []string{"A", "Q"}
	r := map[string]int{}

	for i := 0; i < 19; i++ {
		r[m[rand.Int() % 2]]++
	}
	fmt.Println(r)
}