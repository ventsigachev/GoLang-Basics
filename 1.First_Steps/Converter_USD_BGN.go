//	Write a programm to convert USD to BGN. 1 USD = 1.79549 BGN.


package main

import "fmt"

func main() {

	const i float32 = 1.79549

	var usd float32
	fmt.Scanln(&usd)

	var bgn float32 = usd * i
	fmt.Println(bgn)

}
