package main

import (
		"fmt" 
		"math"
	)

func main() {

	var max int = math.MinInt64
	var min int = math.MaxInt64 
	var n int
	var num int

	fmt.Scanln(&n)

	for i := 1;i <= n;i++ {

		fmt.Scanln(&num)

		if num > max {

			max = num

		}
		if num < min {

			min = num

		}

	}

	fmt.Printf("Max number: %d\n", max)
	fmt.Printf("Min number: %d", min)

}
