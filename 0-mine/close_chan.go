package main

import "fmt"

func new_fun() (stop func(), wasCanceled func() bool) {
	cancel := make(chan struct{})
	wasCanceled = func() bool {
		select {
		case <-cancel:
			return true
		default:
			return false
		}
	}
	stop = func() {
		close(cancel)
	}
	return
}

func main() {
	stop, wasCanceled := new_fun()
	fmt.Println(wasCanceled())
	fmt.Println(wasCanceled())
	stop()
	fmt.Println(wasCanceled())
	fmt.Println(wasCanceled())
}

