package main

import (
	"fmt"
	"math"
)

func main() {

	var a int
	fmt.Scanln(&a)


	for i := 0;i <= a;i++ {
		
		if i % 2 == 0 {

			fmt.Println(math.Pow(2, float64(i)))
		}

	}
}
