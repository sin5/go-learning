package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		12.3,
		15.1,
	}

	fmt.Println(m)

}