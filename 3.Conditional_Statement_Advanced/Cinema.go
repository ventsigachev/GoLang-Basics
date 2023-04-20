// Programm to calculate sum of all tickets.
// Inputs:
// 1.Type of movie(string), 2.Count of rows(int), 3.Count of columns(int).
// Type of movie is: Premiere - cost of ticket = 12.00, Normal = 7.50, Discount = 5.00

package main

import (
	"fmt"
)

func main() {

	var movieType string
	fmt.Scanln(&movieType)

	var rows, cols int
	fmt.Scanln(&rows)
	fmt.Scanln(&cols)

	var allSum float32

	switch movieType {

	case "Premiere":
		allSum = float32(rows) * float32(cols) * 12.00

	case "Normal":
		allSum = float32(rows) * float32(cols) * 7.50

	case "Discount":
		allSum = float32(rows) * float32(cols) * 5.00
	}

	fmt.Printf("%.2f leva", allSum)

}
