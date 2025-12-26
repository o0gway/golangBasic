package main

import "fmt"

const USDToEUR = 0.85
const EURToUSD = 1.18
const RUBToUSD = 79.88
const RUBToEUR = 79.04 / 0.85
const USDToRUB = 0.013
const EURToRUB = 0.011

func main() {
	currency := getCurrency()
	rate1, rate2 := getRates()
	getResult(currency, rate1, rate2)
}

func getCurrency() float64 {
	var currency float64
	for {
		fmt.Print("Пожалуйста введите сумму:")
		_, err := fmt.Scan(&currency)
		if err != nil {
			fmt.Println("Некоретное значение")
			continue
		}
		return currency
	}
}

func getRates() (string, string) {
	var originalRate string
	var targetRate string
	for {
		fmt.Print("Пожалуйста введите исходную валюту(USD/EUR/RUB):")
		fmt.Scanln(&originalRate)
		if originalRate == "RUB" || originalRate == "USD" || originalRate == "EUR" {
			break
		}
		fmt.Println("Некоретное значение")

	}

	for {
		fmt.Print("Пожалуйста введите целевую валюту(USD/EUR/RUB):")
		fmt.Scan(&targetRate)
		if targetRate == "RUB" || targetRate == "USD" || targetRate == "EUR" {
			break
		}
		fmt.Println("Некоретное значение")
	}
	return originalRate, targetRate

}

func getResult(sum float64, originalRate string, targetRate string) {

	if originalRate == "RUB" && targetRate == "USD" {
		fmt.Println(sum * RUBToUSD)
	} else if originalRate == "RUB" && targetRate == "EUR" {
		fmt.Println(sum * RUBToEUR)
	} else if originalRate == "USD" && targetRate == "EUR" {
		fmt.Println(sum * USDToEUR)
	} else if originalRate == "EUR" && targetRate == "USD" {
		fmt.Println(sum * EURToUSD)
	} else if originalRate == "USD" && targetRate == "RUB" {
		fmt.Println(sum * USDToRUB)
	} else if originalRate == "EUR" && targetRate == "RUB" {
		fmt.Println(sum * EURToRUB)
	} else {
		fmt.Println("Нет операций для обмена")
	}
}
