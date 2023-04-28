//	Programm to find if given element is equal the sum of others given elements. 
//	Inputs - sequence of integers.

package main

import (
	"fmt"
	"math"
)

func main() {

	var nums int
	fmt.Scanln(&nums)

	var maxV = math.MinInt64
	var sumAll int = 0
	var v int

	for i := 1; i <= nums; i++ {

		fmt.Scanln(&v)
		sumAll += v

		if v > maxV {

			maxV = v

		}

	}

	var sum int = sumAll - maxV
	if sum == maxV {

		fmt.Printf("Yes\nSum = %d", maxV)

	}else {

		var diff int = int(math.Abs(float64(maxV) - float64(sum)))
		fmt.Printf("No\nDiff = %d", diff)

	}
}
