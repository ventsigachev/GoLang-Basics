package main

import "fmt"

func main() {
	var inch float32
	fmt.Scanln(&inch)

	var cm float32 = inch * 2.54
	fmt.Printf("%.2f", cm)

}
