//	Programm to estimate needed LITERS of water(nw) to fill the aquarium
//	Notice that some percentage of aquarium is filled with sand
//	Note that 1 liter water is equal to 1 cubic decimeter, equal to 0.001 cubic cm.
//	Inputs:
//		f_line - lenght of aquarium(l) in cm
//		s_line - width(w) in cm
//		t_line - height(h) in cm
//		frth_line - sand percentage(p)


package main

import "fmt"

func main() {

	var l int 
	fmt.Scanln(&l)
	var w int
	fmt.Scanln(&w)
	var h int
	fmt.Scanln(&h)
	var sandPercentage float32
	fmt.Scanln(&sandPercentage)
	sandPercentage = sandPercentage / 100

	var volumeAquariumInCm int = h * w * l
	var volumeAquariumInLiters float32 = float32(volumeAquariumInCm) / 1000
	var nw float32 = volumeAquariumInLiters * (1 - float32(sandPercentage))

	fmt.Println(nw)

}
