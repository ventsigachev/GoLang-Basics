//	Programm to calculate amount for painting. 
//	Nylon = 1.5 / m * m
//	Paint = 14.5 / liter
//	Thinner = 5.00 / liter
//	Accessories = 10% more paint, 2.00 m2 nylon, 0.40 for bags, 30% from material expencies for workers per 1 hour.
//	Inputs:
//		f_line - nylon in m2(nM2)
//		s_line - paint in liters(pl)
//		t_line - thinner in liters(tl)
//		fourth_line - working hours(wh)


package main

import "fmt"

func main() {

	var nM2 float32
	fmt.Scanln(&nM2)
	var pl float32
	fmt.Scanln(&pl)
	var morePaint float32 = pl * 0.1
	pl += morePaint
	var tl float32 
	fmt.Scanln(&tl)
	var wh float32
	fmt.Scanln(&wh)

	var nPrice float32 = (nM2 + 2) * 1.5
	var paintPrice float32 = pl * 14.5
	var thinnerPrice float32 = tl * 5
	var allPrice = nPrice + paintPrice + thinnerPrice + 0.4

	var work_price = allPrice * 30 / 100
	var wPrice = work_price * wh

	var endAmount float32 = allPrice + wPrice
	
	fmt.Println(endAmount)

}
