package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name 
}

// func getSum(n1, n2 int) int { can write like this as they are both ints
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(greeting("Harry"))
	fmt.Println(getSum(5, 3))
}