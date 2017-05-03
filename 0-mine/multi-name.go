package main

import "fmt"

func hello() string {
	return "hello world"
}

func main() {
	Hello := hello

	fmt.Println(hello())
	fmt.Println(Hello())
}
