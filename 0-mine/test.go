package main

import "fmt"

type Foo struct {
	A int
	B int
}

func return_p() *Foo {
	return &Foo{1, 2}
}

func return_slice() ([]int, []int) {
	s := []int{1, 2, 3, 4}
	return s, s[2:]
}

func main() {
	a, b := return_slice()
	b[0] = 0
	fmt.Println(a, b)

}