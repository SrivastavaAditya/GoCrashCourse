package main

import "fmt"

func main() {
	fmt.Println(greeting("Aditya"))
	fmt.Println(getSum(3, 6))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}
