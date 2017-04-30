package main

import "fmt"

func main() {
	var s []int
	printSlice("s", s)

	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2)
	printSlice("s", s)

	s = append(s, 3, 4, 5, 6, 7)
	printSlice("s", s)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}