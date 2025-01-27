package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

type Temperature struct {
	day1 float64
	day2 float64
}

func main() {
	var productNames [4]string
	productNames = [4]string{"A Book"}

	productNames[2] = "A Carpet"

	prices := [4]float64{10.99, 9.99, 45.99, 20}

	fmt.Println(prices)
	fmt.Println(productNames)

	featuredPrices := prices[1:]
	fmt.Println(featuredPrices)
	higlightedPrices := featuredPrices[:1]
	fmt.Println(higlightedPrices)

	fmt.Println(prices)
	featuredPrices[0] = 199.99
	fmt.Println(prices)

	fmt.Println(prices)
	higlightedPrices[0] = 99.99
	fmt.Println(prices)

	fmt.Println(len(featuredPrices), cap(featuredPrices))

	fmt.Println(len(higlightedPrices), cap(higlightedPrices))
}
