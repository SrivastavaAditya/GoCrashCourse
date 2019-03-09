package main

import (
	"fmt"
	"math")

//Defining Interface
type Shape interface{
	area() float64
}

type Circle struct{
	x, y, radius float64
}

type Rectangle struct{
	width, height float64
}

func (c Circle) area() float64{
	return math.Pi*c.radius*c.radius
}

func (r Rectangle) area() float64{
	return r.width*r.height
}

func getArea(s Shape) float64{
	return s.area()
}

func main(){
	c := Circle{0, 0, 5}
	r := Rectangle{10, 15}

	fmt.Printf("Circle %f\n", getArea(c))
	fmt.Printf("Rectangle %f\n", getArea(r))
}