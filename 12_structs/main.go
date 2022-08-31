package main

import (
	"fmt"
	"strconv"
)

//Define Struct
type User struct {
	id 			int
	name 		string
	workplace 	string
	age 		int
}

// Getting method VALUE RECIEVER
// USING STRING CONVERTER AS YOU CAN NOT HAVE MISMATCHED TYPES
func(user User) greet() string {
	return "Hello, my name is " + user.name + " I am " + strconv.Itoa(user.age)
}
// hasBirthday method POINTER RECIEVER (changes data)
func(user *User) hasBirthday() {
	user.age++
}
// Change workplace POINTER RECIEVER
func(user *User) updateWorkplace(workplace string) {
	user.workplace = workplace
}

func main() {

	//Int User via struct
	user1 := User{id: 2, name: "Harry", workplace: "nology", age: 26}
	user2 := User{3, "Danni", "twitch", 27}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user1.name)
	user1.age++
	fmt.Println(user1)

	fmt.Println(user1.greet())
	user1.hasBirthday()
	fmt.Println(user1.age)
	user1.updateWorkplace("TrustGrid")
	fmt.Println(user1.workplace)
}