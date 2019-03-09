package main

import "fmt"

func main() {
	//Define Map
	//emails := make(map[string]string)

	//Assign KV
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	//Declare map and add KV
	emails := map[string]string{"Bob": "bob@gmail.com", "Mike": "mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	//Delete from Map
	// delete(emails, "Bob")
	// fmt.Println(emails)
}
