package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	collectInput("Revenue: ", &revenue)
	//fmt.Print("Revenue: ")
	//fmt.Scan(&revenue)

	collectInput("Expenses: ", &expenses)
	//fmt.Print("Expenses: ")
	//fmt.Scan(&expenses)

	collectInput("Tax Rate: ", &taxRate)
	//fmt.Print("Tax Rate: ")
	//fmt.Scan(&taxRate)

	earningsBeforeTax, profit, ratio := calculateProfit(revenue, expenses, taxRate)
	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	// ratio := earningsBeforeTax / profit

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func collectInput(text string, value *float64) {
	fmt.Print(text)
	fmt.Scan(value)
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt / profit

	return ebt, profit, ratio
}
