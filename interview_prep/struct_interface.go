// Struct and interface : Create a Shape interface with methods Perimeter and Area. Implement it for Circle and Rectangle structs.

package main

import (
	"fmt"
	"math"
)

// Shape interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	radius float64
}

func (x Circle) Area() float64      { return math.Pi * x.radius * x.radius }
func (x Circle) Perimeter() float64 { return 2 * math.Pi * x.radius }

// Rectangle struct
type Rectangle struct {
	length float64
	width  float64
}

func (x Rectangle) Area() float64      { return x.length * x.width }
func (x Rectangle) Perimeter() float64 { return 2 * (x.length + x.width) }

func main() {
	//
	myCircle := Circle{
		radius: 3.56,
	}
	fmt.Println("circle area: ", myCircle.Area())
	fmt.Println("circle perimiter: ", myCircle.Perimeter())

	//
	myRectangle := Rectangle{
		length: 4,
		width:  2,
	}
	fmt.Println("rect area: ", myRectangle.Area())
	fmt.Println("rect perimiter: ", myRectangle.Perimeter())
}
