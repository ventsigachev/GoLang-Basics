//	Programm to find percentage of int in range < 200(p1), between 200 - 399(p2), int[400 - 599](p3),
//	int[600 - 799](p4), int[ >= 800](p5).

package main

import "fmt"

func main() {

	var nums int
	fmt.Scanln(&nums)

	var p1Counter, p2Counter, p3Counter, p4Counter, p5Counter int = 0, 0, 0, 0, 0

	for i := 1; i <= nums; i++ {

		var n int
		fmt.Scanln(&n)

		if n < 200 {

			p1Counter++

		} else if n >= 200 && n <= 399 {

			p2Counter++

		} else if n >= 400 && n <= 599 {

			p3Counter++

		} else if n >= 600 && n <= 799 {
			
			p4Counter++

		} else {

			p5Counter++

		}

	}

	var p1 float32 = float32(p1Counter) / float32(nums) * 100
	var p2 float32 = float32(p2Counter) / float32(nums) * 100
	var p3 float32 = float32(p3Counter) / float32(nums) * 100
	var p4 float32 = float32(p4Counter) / float32(nums) * 100
	var p5 float32 = float32(p5Counter) / float32(nums) * 100

	fmt.Printf("%.2f%%\n%.2f%%\n%.2f%%\n%.2f%%\n%.2f%%", p1, p2, p3, p4, p5)
	
}
