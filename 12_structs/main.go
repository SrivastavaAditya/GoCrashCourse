package main

import (
	"fmt"
	"strconv")

type Person struct{
	firstName string
	lastName string
	city string
	gender string
	age int
}

//Greeting method (value receiver)
func greeting(p Person) string{
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and my age is " + strconv.Itoa(p.age)
}

//has Birthday method (pointer receiver)
func (p *Person) hasBirthday(){
	p.age++
}

//get married
func (p *Person) getMarried(spouseLastName string){
	if p.gender == "m" || p.gender == "M"{
		return
	}else{
		p.lastName = spouseLastName
	}
}


func main(){
	// Init Person jusing Struct
	person1 := Person {
		firstName: "Samantha",
		lastName: "Smith",
		city: "Boston",
		gender: "f",
		age: 25}

	// Alternative
	person2 := Person{"Aditya", "Srivastava", "Noida", "M", 24}
	fmt.Println(person2)
	person2.age++
	fmt.Println(greeting(person1))
	person2.hasBirthday()
	fmt.Println(greeting(person2))
	person2.getMarried("Katyal")
	person1.getMarried("Katyal")
	fmt.Println(person1, person2)
}