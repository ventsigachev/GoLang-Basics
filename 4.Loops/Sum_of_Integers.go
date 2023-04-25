package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	var sum int
	var m int

	for i := 1;i <= n;i++ {
		
		fmt.Scanln(&m)
		sum += m
		
	}

	fmt.Println(sum)

}
