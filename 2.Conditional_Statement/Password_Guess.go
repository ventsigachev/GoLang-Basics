// Write a programm which checked if the entered password is equal to "s3cr3t!P@ssw0rd".
// If so print the message "Wellcome", if not - "Wrong Password".

package main

import "fmt"

func main() {

	const pass string = "s3cr3t!P@ssw0rd"

	var p string
	fmt.Scanln(&p)

	if p == pass {

		fmt.Println("Welcome")

	} else {

		fmt.Println("Wrong password!")

	}
}
