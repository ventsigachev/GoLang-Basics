//	Programm to estimate cost of equipment for basketball training.
//	Basketball sneakers(bs) - 40% from year fee
//	Basketball teeshurt(bt) - 20% from Basketball sneakers
//	Basketball ball(bb) - 1/4 from price of Basketball teeshurt
//	Bssketball accessories(ba) - 1/5 from price of Basketball ball
//	Input:
//		year fee(yf)


package main

import "fmt"

func main() {

	var yf float32
	fmt.Scanln(&yf)

	var bs float32 = yf - (yf * 40 / 100)
	var bt float32 = bs - (bs * 20 / 100)
	var bb float32 = bt / 4
	var ba float32 = bb / 5

	var cost float32 = bs + bt + bb + ba + yf
	fmt.Println(cost)

}