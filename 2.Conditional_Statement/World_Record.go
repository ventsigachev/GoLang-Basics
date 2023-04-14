// Improving world swimm record. Water resistance 12.5 sec for every 15 meters.
// Inputs:
//		f_line - world record
//		s_line - distance in meters
//		t_line - time in seconds for 1 meter swimming

package main

import (
	"fmt"
	"math"
)

func main() {

	var wRecord, dist, timeM float64
	fmt.Scanln(&wRecord)
	fmt.Scanln(&dist)
	fmt.Scanln(&timeM)

	var allTime float64 = dist * timeM

	if dist >= 15 {

		var resistance float64 = 12.5 * math.Floor(dist / 15)
		allTime += resistance

	}

	var rest float64 = allTime - wRecord

	if allTime < wRecord {

		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", allTime)

	} else {

		fmt.Printf("No, he failed! He was %.2f seconds slower.", rest)
	}
}
