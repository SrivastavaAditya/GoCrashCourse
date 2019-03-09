package main

import "fmt"

func main() {
	fmt.Println(greeting("Aditya"))
	fmt.Println(getSum(3, 6))
	fmt.Println(getDifference(23, 987))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func getDifference(num1 int, num2 int) int{
	if num1>=num2{
		return num1-num2
	}else{
		return num2-num1
	}
}
