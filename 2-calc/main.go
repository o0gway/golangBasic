package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	operation := getOperation()
	numbers := make([]int, 0, 1)
	for {
		numbers = getNumbers()
		if numbers != nil {
			break
		}
	}
	makeCalculate(operation, numbers)
}

func getOperation() string {
	var userChoice string
	for {
		fmt.Print("Пожалуйста введите команду для расчетов(AVG/SUM/MED): ")
		fmt.Scan(&userChoice)
		userChoice = strings.ToUpper(userChoice)
		if userChoice == "AVG" || userChoice == "SUM" || userChoice == "MED" {
			return userChoice
		}
		fmt.Println("Неверная команда")
	}
}

func getNumbers() []int {
	var userNumbers string
	fmt.Print("Введите цифры для расчетов через запятую: ")
	fmt.Scanln(&userNumbers)
	userNumbers = strings.ReplaceAll(userNumbers, " ", "")
	numbersSliceStr := strings.Split(userNumbers, ",")
	numbersSliceInt := make([]int, 0, len(numbersSliceStr))
	for _, value := range numbersSliceStr {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("В списке ваших значений есть не цифровое значение")
			return nil
		}
		numbersSliceInt = append(numbersSliceInt, num)
	}
	return numbersSliceInt
}

func makeCalculate(opr string, nums []int) {
	if opr == "AVG" {
		var sum int
		for _, value := range nums {
			sum += value
		}
		avg := sum / len(nums)
		fmt.Printf("Среднее значение: %d\n", avg)
	} else if opr == "SUM" {
		var sum int
		for _, value := range nums {
			sum += value
		}
		fmt.Printf("Сумма равна: %d\n", sum)
	} else if opr == "MED" {
		nums = sortNums(nums)
		if len(nums)%2 != 0 {
			med := (len(nums) / 2) + 1
			fmt.Printf("Медианное значение: %d\n", med)
		} else {
			num1 := nums[len(nums)/2] - 1
			num2 := nums[len(nums)/2]
			med :=  float64(num1 + num2) / 2.0
			fmt.Printf("Медианное значение: %v\n", med)
		}
	}
}

func sortNums(nums []int) []int {
	for index := range nums {
		_ = index
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}
