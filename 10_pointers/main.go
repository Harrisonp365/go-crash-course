package main

import "fmt"

func main() {
	a := 5
	b := &a 

	fmt.Println(a)
	fmt.Println(b) // points to memory address of A
	fmt.Printf("%T\n", b)
	//Use * to read value from address
	fmt.Println(*b)

	//Change value with pointer
	*b = 10
	fmt.Println(a)
}