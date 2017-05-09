package main

import "fmt"

type A struct {
}

type C struct {

}

type B struct {
	A
	C
}

func (a *A)in_A_B() string {
	return "A"
}

func (b *B)in_A_B() string {
	return "B"
}

func (a *A)in_A_C() string {
	return "A"
}

func (c *C)in_A_C() string {
	return "C"
}

func (a *A)in_A() string {
	return "A"
}

func main() {
	b := B{}
	fmt.Println("b.in_A_B()", b.in_A_B())
	fmt.Println("b.in_A()", b.in_A())
	fmt.Println("b.A.in_A()", b.A.in_A())
	fmt.Println("b.A.in_A_C()", b.A.in_A_C())
	fmt.Println("b.C.in_A_C()", b.C.in_A_C())
}