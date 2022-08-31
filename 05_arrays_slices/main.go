package main

import "fmt"

func main() {
	//Arrays fixed lengths and types
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Kiwi"

	fmt.Println(fruitArr)

	//Declare and assign
	fruits := [2]string{"Strawberry", "grape"}
	fmt.Println(fruits)

	//Slices
	//Slices can change size unlike an array
	fruitSlice := []string{"apple", "kiwi", "grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}