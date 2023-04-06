//	Programm to estimate costs for food delivery
//	Chicken menu is 10.35, fish menu is 12.40, vegan menu - 8.15, desert = 20% of food cost, delivery is 2.50
//	Inputs:
//		f_line - count of chicken menus
//		s_line - count of fish menus
//		t_line - count of vegan menus


package main

import "fmt"

func main() {

	var cm float32
	fmt.Scanln(&cm)
	var fm float32
	fmt.Scanln(&fm)
	var vm float32
	fmt.Scanln(&vm)
	var delivery float32 = 2.50

	var cmPrice float32 = cm * 10.35
	var fmPrice float32 = fm * 12.40
	var vmPrice float32 = vm * 8.15
	var desert float32 = (cmPrice + fmPrice + vmPrice) * 20 / 100
	

	var allCosts float32 = cmPrice + fmPrice + vmPrice + desert + delivery
	fmt.Println(allCosts)

}
