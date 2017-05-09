package main

import "fmt"

type Channel struct {
	in  <-chan int
	out chan <- int
}

func main() {
	c := &Channel{}
	ch := make(chan int, 1)
	c.in = ch
	c.out = ch

	c.out <- 2

	fmt.Println(<-c.in)
}