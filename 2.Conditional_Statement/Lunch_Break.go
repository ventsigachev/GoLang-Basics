// You need to evaluate if you have time for movie and lunch in your lunch break.
// Time for lunch is 1/8 of break time and time for rest is 1/4 of break time.
// Inputs:
//	1. Movie name, 2. Movie duration, 3. Total break time.


package main

import "fmt"
import "math"

func main () {
	
	var seriesName string
	fmt.Scanln(&seriesName)

	var episodeDuration int
	fmt.Scanln(&episodeDuration)

	var totalTimeBreak int
	fmt.Scanln(&totalTimeBreak)

	var timeLunch float32 = float32(totalTimeBreak) / 8.0 
	var timeRest float32 = float32(totalTimeBreak) / 4.0
	var timeSeries float32 = float32(totalTimeBreak) - timeLunch - timeRest

	if timeSeries >= float32(episodeDuration) {

		var leftTime float32 = timeSeries - float32(episodeDuration) 
		fmt.Printf("You have enough time to watch %s and left with %.0f minutes free time.", seriesName, math.Ceil(float64(leftTime)))

	} else {
		
		var needTime float32 = float32(episodeDuration) - timeSeries 
		fmt.Printf("You don't have enough time to watch %s, you need %.0f more minutes.", seriesName, math.Ceil(float64(needTime)))

	}

}
