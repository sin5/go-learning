package main

import "fmt"

var checkErr error

func main() {
	if checkErr == nil {
		fmt.Println(checkErr)
	}
}