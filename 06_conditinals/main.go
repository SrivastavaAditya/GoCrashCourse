package main

import "fmt"

func main() {
	//If
	// x := 15
	// y := 10

	// if x < y {
	// 	fmt.Printf("%d is less than %d", x, y)
	// } else {
	// 	fmt.Printf("%d is less than %d", y, x)
	// }

	//Else if
	// color := "blue"

	// if color == "red" {
	// 	fmt.Printf("Color is red")
	// } else if color == "blue" {
	// 	fmt.Printf("Color is blue")
	// } else {
	// 	fmt.Printf("Color not defined")
	// }

	//Switch
	color := "blue"
	switch color {
	case "red":
		fmt.Printf("Color is red")

	case "blue":
		fmt.Printf("color is blue")

	default:
		fmt.Printf("Color not defined")
	}
}
