//	Write a programm, which reads a number, and prints all the int to 1.

package main

import "fmt"

func main() {

	var a int
	fmt.Scanln(&a)


	for i := a; i >= 1; i-- {
		
		fmt.Println(i)

	}
}
