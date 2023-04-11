package main

import (
	"fmt"
	"math"
)

func main() {

	var figure string
	fmt.Scanln(&figure)

	if figure == "square" {

		var side float32
		fmt.Scanln(&side)
		fmt.Printf("%.3f", side * side)

	} else if figure == "rectangle" {

		var a, b float32
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		fmt.Printf("%.3f", a * b)
		
	} else if figure == "circle" {

		var r float32
		fmt.Scanln(&r)
		fmt.Printf("%.3f", math.Pi * r * r)

	} else if figure == "triangle" {

		var a, h float32
		fmt.Scanln(&a)
		fmt.Scanln(&h)
		fmt.Printf("%.3f", (a * h) / 2)

	}

}
