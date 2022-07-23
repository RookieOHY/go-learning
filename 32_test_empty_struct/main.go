package main

import "fmt"

type EmptyStruct struct {
}

func main() {
	a := struct{}{}
	b := struct{}{}
	c := EmptyStruct{}

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	fmt.Println("test")
}
