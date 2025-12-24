package main

import "fmt"

func main() {

	const USDToEUR = 0.85
	const RUBToUSD = 79.88
	const EURToRub = 79.88 / 0.85

}

func getCurrency() float64 {
	var currency float64
	fmt.Scan(&currency)
	return currency
}

func getRate(currency float64, rate1 string, rate2 string) {

}
