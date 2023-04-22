//Inputs:
//	1. and 2. Integers
//	3. Operator - one of "+", "-", "*", "/", "%".
//Outputs:
//	1. One line:
//		If "+", "-", "*" - "{N1} {operator} {N2} = {result} – {even/odd}".
//		If "/" - "{N1} / {N2} = {result}".
//		If "%" - "{N1} % {N2} = {remnant}".
//		If zero division - "Cannot divide {N1} by zero".


package main

import "fmt"

func main () {

	var n1, n2 int
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)

	var operator string
	fmt.Scanln(&operator)


	switch operator {
		case "+":
		
			var sum int = n1 + n2
			if sum % 2 == 0 {
				
				fmt.Printf("%d + %d = %d - even", n1, n2, sum)
			} else {
				
				fmt.Printf("%d + %d = %d - odd", n1, n2, sum)
			}
		case "-":
			
			var diff int = n1 - n2
			if diff % 2 == 0 {
			
				fmt.Printf("%d - %d = %d - even", n1, n2, diff)
			} else {

				fmt.Printf("%d - %d = %d - odd", n1, n2, diff)
			}
		case "*":

			var result int = n1 * n2
			if result % 2 == 0 {

				fmt.Printf("%d * %d = %d - even", n1, n2, result)
			} else {

				fmt.Printf("%d * %d = %d - odd", n1, n2, result)
			}
		case "/":

			if n2 == 0 {
				fmt.Printf("Cannot divide %d by zero", n1)

			} else {
			
				var division float32 = float32(n1) / float32(n2)
				fmt.Printf("%d / %d = %.2f", n1, n2, division)
			}
	
		case "%":

			if n2 == 0 {
				fmt.Printf("Cannot divide %d by zero", n1)

			} else {

				var reminder int = n1 % n2
				fmt.Printf("%d %% %d = %d", n1, n2, reminder)
			}
	}

}
