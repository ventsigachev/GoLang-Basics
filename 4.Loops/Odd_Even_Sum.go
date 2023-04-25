package main

import (
	"fmt"
	"math"
)

func main() {

	var sumOdd int
	var sumEven int
	var v int
	var n int

	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {

		fmt.Scanln(&v)
		if i % 2 == 0 {

			sumEven += v

		} else {

			sumOdd += v

		}
		

	}

	
	if sumEven == sumOdd {

		fmt.Printf("Yes\nSum = %d", sumEven)

	}else {

		var diff int = int(math.Abs(float64(sumEven) - float64(sumOdd)))
		fmt.Printf("No\nDiff = %d", diff)

	}

}
