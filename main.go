package main

import (
	"errors"
	"fmt"
	"math"
)
const IMTPower = 2
func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела равняется: %.0f ", IMT)
	fmt.Print(result)
	islean := IMT < 16
	switch {
	case islean:
		fmt.Println("Сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("Дефицит массы тела")
	case IMT < 25:
		fmt.Println("Норма")
	case IMT < 30:
		fmt.Println("Избыточная масса")
	case IMT < 35:
		fmt.Println("1-я степень ожирения")
	default:
		fmt.Println("2-я степень ожирения")
	}
}
func calculateIMT(userHeight, userMass float64) (float64, error) {
	if userHeight <= 0 || userMass <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userMass / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}
func GetUserInput() (float64, float64) {
	var userHeight float64
	var userMass float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userMass)
	return userHeight, userMass
}
func Repeat() bool {
	var userChoise string
	fmt.Print("Хотите продолжить (Да/Нет)? ")
	fmt.Scan(&userChoise)
	if userChoise == "Да" || userChoise == "да" {
		return true
	}
	return false
}
func main() {
	for {
		fmt.Println("___Калькулятор ИМТ___")
		userHeight, userMass := GetUserInput()
		IMT, err := calculateIMT(userHeight, userMass)
		if err != nil {
			fmt.Println("Заданы не корректные значения")
			continue
		}
		outputResult(IMT)
		if Repeat() {
			continue
		}
		break
	}
}
