package main

import "fmt"


var rateMaps map[string]float64 = map[string]float64{
	"EUR": 93.55,
	"USD": 79.45,
	"RUB": 1,
}


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
			fmt.Println("Некорректное значение")
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
		fmt.Println("Некорректное значение")

	}

	for {
		fmt.Print("Пожалуйста введите целевую валюту(USD/EUR/RUB):")
		fmt.Scan(&targetRate)
		if targetRate == "RUB" || targetRate == "USD" || targetRate == "EUR" {
			break
		}
		fmt.Println("Некорректное значение")
	}
	return originalRate, targetRate

}

func getResult(sum float64, originalRate string, targetRate string) {

	if originalRate == "RUB" && targetRate == "USD" {
		fmt.Println(sum * (rateMaps[targetRate] / rateMaps[originalRate]))
	} else if originalRate == "RUB" && targetRate == "EUR" {
		fmt.Println(sum * (rateMaps[targetRate] / rateMaps[originalRate]))
	} else if originalRate == "USD" && targetRate == "EUR" {
		fmt.Println(sum * (rateMaps[originalRate] / rateMaps[targetRate]))
	} else if originalRate == "EUR" && targetRate == "USD" {
		fmt.Println(sum * (rateMaps[targetRate] / rateMaps[originalRate]))
	} else if originalRate == "USD" && targetRate == "RUB" {
		fmt.Println(sum * rateMaps[originalRate])
	} else if originalRate == "EUR" && targetRate == "RUB" {
		fmt.Println(sum * rateMaps[originalRate])
	} else {
		fmt.Println("Нет операций для обмена")
	}
}
