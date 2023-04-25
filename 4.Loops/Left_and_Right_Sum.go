package main

import (
	"fmt"
	"math"
)

func main() {

	var sum1 int
	var sum2 int
	var v int
	var n int

	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {

		fmt.Scanln(&v)
		sum1 += v

	}

	for i := 1; i <= n; i++ {

		fmt.Scanln(&v)
		sum2 += v
	}

	if sum1 == sum2 {

		fmt.Printf("Yes, sum = %d", sum1)

	}else {

		var diff int = int(math.Abs(float64(sum1) - float64(sum2)))
		fmt.Printf("No, diff = %d", diff)

	}

}
