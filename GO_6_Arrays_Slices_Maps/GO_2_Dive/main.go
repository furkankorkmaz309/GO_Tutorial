package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(len(prices), cap(prices))
	prices[1] = 9.99
	//  prices[2] = 11.99 WONT WORK

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

	prices = append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

	prices = prices[1:]
	fmt.Println(prices)

	prices2 := []float64{10.99, 8.99}
	fmt.Println(prices2)

	prices2 = append(prices2, 5.99, 12.99, 29.99, 100.50)
	fmt.Println(prices2)

	discountPrices := []float64{101.99, 80.99, 20.50}
	prices2 = append(prices2, discountPrices...)
	fmt.Println(prices2)
}
