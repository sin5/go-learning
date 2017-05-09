package main

import "fmt"

func main() {

	var bufa [64]byte
	buf := bufa[:]
	b := []byte{}
	b = append(b, buf[:2]...)
	b = append(b, buf[:2]...)
	fmt.Println(b)
}
