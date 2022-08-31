package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	//Assign key vals
	emails["Harry"] = "harry@gmail.com"
	emails["Bob"] = "bob@gmail.com"
	emails["Danni"] = "danni@gmail.com"
	emails["Laura"] = "laura@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Harry"]) //Where you would use an ID in real world APP

	//Delete from map
	delete(emails, "Harry")
	fmt.Println(emails)

	//Declare and assign 

	users := map[string]string{"1": "Harry", "2": "Bob", "3": "Danni"}
	fmt.Println(users)
	users["4"] = "Laura"
	fmt.Println(users)
	delete(users, "4")
	fmt.Println(users)
}