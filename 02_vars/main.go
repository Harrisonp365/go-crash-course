package main

import "fmt"

func main() {
	//Using Var is not CONST, replace var with const if const is wanted
	//Don't have to write STRING or INT as it is inferred
	var name string = "Harry"
	var age int = 26
	//casting to a type
	var age32 int32 = 38
	var isGo = true

	fmt.Println(name, age)
	fmt.Printf("%T\n", age32)
	fmt.Printf("%T\n", isGo)

	//Shorthand variable decleration
	shortHandName := "Harrison"
	size := 1.4

	//multiple decleration
	height, weight := 186, 75

	fmt.Println(shortHandName)
	fmt.Printf("%T\n", size)
	fmt.Println(height, weight)

}