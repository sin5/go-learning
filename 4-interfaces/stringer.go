package main

import "fmt"

type Person struct {
	Name string
	Age  int8
}

func (p *Person)String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	w := &Person{"WuXin", 25}
	fmt.Println(w)
}