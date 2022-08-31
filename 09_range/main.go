package main

import "fmt"

func main() {
	ids := []int{1,2,3,4,5,6,7,8,9,10}

	//Loops through ids
	for i, id := range ids {
		fmt.Printf("%d index - ID: %d\n", i, id)
	}

	//not using index UNDERSCORE IS FOR UNUSED PARAM
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum of IDs", sum)

	// Range with map
	users := map[string]string{"1": "Harry", "2": "Bob", "3": "Danni"}

	for k, v := range users {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range users {
		fmt.Println("Key: " + k)
	}
}