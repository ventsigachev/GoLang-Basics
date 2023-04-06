//	Programm to calculate overall amount for school necessities.
//	Packet of pens = 5.80(pp)
//	Packet of markers = 7.20 (pm)
//	Liter of chemical = 1.20 (lc)
//	Inputs:
//		fitst line - pp
//		second line - pm
//		third line - lc
//		fourth line - discount (d) in %


package main

import "fmt"

func main() {

	var pp float32 
	fmt.Scanln(&pp)
	var pm float32
	fmt.Scanln(&pm)
	var lc float32
	fmt.Scanln(&lc)
	var d float32
	fmt.Scanln(&d)
	d /= 100


	var sum float32 = pp * 5.8 + pm * 7.2 + lc * 1.2
	var dSum = sum * d
	var endSum = sum - dSum

	fmt.Println(endSum)

}
