// Programm to convert radians in degrees. 1 degree = 1 radian * 180 / π.

package main

import (
	"fmt";
	"math"
)

func main() {
	const i float32 = 180 / math.Pi

	var rad float32
	fmt.Scanln(&rad)

	var grad float32 = rad * i
	fmt.Println(grad)

}
