// Write a programm, which reads an int, and prints "invalid", if number != 0 and != [100 ... 200].

package main 

import "fmt"

func main() {

	var i int
	fmt.Scanln(&i)

	if i != 0 && (i < 100 || i > 200) {

		fmt.Println("invalid")

	}
}
