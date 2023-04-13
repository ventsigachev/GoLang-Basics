// Three athlets finished for amount of seconds. Calculate the sum of time of the athlets as minutes:seconds.
// The seconds must have leading zero if nesessary.

package main

import "fmt"

func main() {

	var t1, t2, t3 int
	fmt.Scanln(&t1)
	fmt.Scanln(&t2)
	fmt.Scanln(&t3)

	var sum int = t1 + t2 + t3
	var min int = sum / 60
	var sec int = sum % 60

	if sec < 10 {

		fmt.Printf("%d:0%d", min, sec)

	} else {

		fmt.Printf("%d:%d", min, sec)

	}

}
