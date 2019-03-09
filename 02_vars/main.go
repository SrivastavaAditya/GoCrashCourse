package main

import "fmt"

func main() {
	// Using var keyword
	// var name = "Aditya"
	var age int32 = 24
	var isCool = true
	isCool = false

	//Short Hand Method
	name, email := "Aditya", "aditya@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
}
