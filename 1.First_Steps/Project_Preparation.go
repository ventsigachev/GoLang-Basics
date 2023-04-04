// Create a programm to calculate how many hours will take to an architect, to finish all projects,
// if it takes 3 hour for one project.
// Inputs:
//		First line - architect's name
//		Second line - projects' count

package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)
	var counts int
	fmt.Scanln(&counts)

	var time int = 3 * counts

	fmt.Printf("The architect %s will need %d hours to complete %d project/s.", name, time, counts)
}
