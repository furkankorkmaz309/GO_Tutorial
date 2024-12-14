package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue : ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses : ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate : ")
	fmt.Scan(&taxRate)

	var earningsBeforeTax float64 = revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)

	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings Before Tax is", earningsBeforeTax)
	fmt.Println("Earnings After Tax (profit) is", earningsAfterTax)
	fmt.Println("Ratio is", ratio)

}
