package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 3}
	p := &v
	p.X = 4
	fmt.Println(v)
}