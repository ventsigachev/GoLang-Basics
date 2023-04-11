// Input:
// 		First line - Text - one of the following square, rectangle, circle или triangle.
//		Nest line(s) 	- if text is square = side of square - float
//						- if text is rectangle = sides of rectangle - floats
//						- if text is circle = radius of the cirlce - float
//						- if text is trialngle = side and height of the triangle - floats
// Output:
// 		Calculated area of the figure, rounded to 3 digits after the comma.

package main

import (
	"fmt"
	"math"
)

func main() {

	var area float32
	var pi float32 = math.Pi
	
	var figure string
	fmt.Scanln(&figure)

	if figure == "square" {

		var side float32
		fmt.Scanln(&side)
		area  = side * side
		fmt.Printf("%.3f", area)

	} else if figure == "rectangle" {

		var a, b float32
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		area = a * b
		fmt.Printf("%.3f", area)
		
	} else if figure == "circle" {

		var r float32
		fmt.Scanln(&r)
		area = pi * r * r
		fmt.Printf("%.3f", area)

	} else if figure == "triangle" {

		var a, h float32
		fmt.Scanln(&a)
		fmt.Scanln(&h)
		area = (a * h) / 2
		fmt.Printf("%.3f", area)

	}

}
