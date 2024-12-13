package main

import "fmt"

var revenue float64
var expenses float64
var taxRate float64

func main() {
	revenue, expenses, taxRate = inputs(revenue, expenses, taxRate)
	calculations(revenue, expenses, taxRate)

}

func inputs(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	fmt.Print("Revenue : ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses : ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate : ")
	fmt.Scan(&taxRate)

	return revenue, expenses, taxRate
}

func calculations(revenue float64, expenses float64, taxRate float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings Before Tax is", earningsBeforeTax)
	fmt.Println("Earnings After Tax (profit) is", earningsAfterTax)
	fmt.Println("Ratio is", ratio)
}
