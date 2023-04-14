// Inputs:
// 		first_line - hours(0 - 23)
// 		sec_line - seconds(0 - 59)
// Output:
// 		Time + 15 minutes in manner hours:minutes, with leading zero of minutes if needed.


package main

import "fmt"

func main() {

	var h int
	var m int
	fmt.Scanln(&h)
	fmt.Scanln(&m)

	m += 15

	if m == 60 {

		h += 1
		m = 0

	}else if m > 60 {

		h += 1
		m -= 60

	}

	if h == 24 {

		h = 0
	}

	if m < 10 {

		fmt.Printf("%d:0%d", h, m)

	}else {

		fmt.Printf("%d:%d", h, m)

	}

}
