package main

import "fmt"

func main() {
	var v [2]string
	v[0] = "Hello"
	v[1] = "World"
	fmt.Println(v[0], v[1])
	fmt.Println(v)
}