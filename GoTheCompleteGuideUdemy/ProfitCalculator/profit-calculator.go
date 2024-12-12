package main

import (
	"errors"
	"fmt"
	"os"
)

const ResultFileName = "results.txt"

func main() {
	revenue, errorRevenue := getUserInput("Revenue: ")
	expenses, errorExpenses := getUserInput("Expenses: ")
	taxRate, errorTaxRate := getUserInput("Tax Rate: ")

	if errorRevenue != nil || errorExpenses != nil || errorTaxRate != nil {
		fmt.Println("Value must be a positive number")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	os.WriteFile(ResultFileName, []byte(fmt.Sprintf("EBT: %.1f, Profit: %.1f, Ratio: %.3f", ebt, profit, ratio)), 0644)
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid input")
	}
	return userInput, nil
}
