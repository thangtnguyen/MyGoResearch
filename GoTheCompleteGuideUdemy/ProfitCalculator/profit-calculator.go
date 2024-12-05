package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio := earningsBeforeTax / profit

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
