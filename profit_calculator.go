package main

import "fmt"

func main() {
	var revernue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revernue: ")
	fmt.Scan(&revernue)

	fmt.Print("expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revernue - expenses

	profit := float64(ebt) * (1 - taxRate/100)

	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
