package main

import (
	"fmt"
	"math"
)

func main() {
	// SOLUTION 1
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	//investmentAmount and years variables should not be written as int, they should be converted to float64
	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println("Solution 1 :", futureValue)

	//
	// SOLUTION 2
	var investmentAmount2 float64 = 1000
	var expectedReturnRate2 float64 = 5.5
	var years2 float64 = 10

	var futureValue2 = investmentAmount2 * math.Pow((1+expectedReturnRate2/100), years2)
	fmt.Println("Solution 2 :", futureValue2)

	//
	// SOLUTION 3
	investmentAmount3, years3, expectedReturnRate3 := 1000.0, 10.0, 5.5

	var futureValue3 = investmentAmount3 * math.Pow((1+expectedReturnRate3/100), years3)
	fmt.Println("Solution 3 :", futureValue3)

	//
	// PART 2
	fmt.Println("")
	fmt.Println("*** PART 2 ***")

	const inflationRate float64 = 2.5
	futureRealValue := futureValue3 / math.Pow(1+inflationRate/100, years3)

	fmt.Println(futureRealValue)

	//
	// PART 3
	fmt.Println("")
	fmt.Println("*** PART 3 ***")

	const inflationRate4 = 6.5
	var investmentAmount4 float64 = 1000
	// OR
	// var investmentAmount4 float64
	years4 := 10.0
	expectedReturnRate4 := 5.5

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount4)
	// If const exists, scan operation cannot be applied on it

	fmt.Print("Years : ")
	fmt.Scan(&years4)

	fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate4)

	futureValue4 := investmentAmount4 * math.Pow(1+expectedReturnRate4/100, years4)
	futureRealValue4 := futureValue4 / math.Pow(1+inflationRate4/100, years4)

	fmt.Println(futureValue4)
	fmt.Println(futureRealValue4)

	//
	// PART 4
	fmt.Println("")
	fmt.Println("*** PART 4 - Formatting ***")
	fmt.Printf("Future Value: %v\n", futureValue4)
	fmt.Printf("Future Value (adjusted for Inflation): %v\n", futureRealValue4)
	fmt.Println("")

	fmt.Printf("Future Value: %f\n", futureValue4)
	fmt.Printf("Future Value (adjusted for Inflation): %f\n", futureRealValue4)
	fmt.Println("")

	fmt.Printf("Future Value: %.0f\n", futureValue4)
	fmt.Printf("Future Value (adjusted for Inflation): %.0f\n", futureRealValue4)

}
