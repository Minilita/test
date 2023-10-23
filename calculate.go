package main

import (
	"fmt"
	"os"
)

func main() {
	// Создаем словарь
	var m = map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	//Создаём переменные для хранения ввода
	var first string
	var second string
	var operation string

	//Читаем 1 число и записываем в переменную
	fmt.Print("Введите первое число: ")
	fmt.Fscan(os.Stdin, &first)
	var firstInt = m[first]
	if firstInt == 0 {
		fmt.Print("Ошибка")
	} else {
		//Читаем 2 число и записываем в переменную
		fmt.Print("Введите второе число: ")
		fmt.Fscan(os.Stdin, &second)
		var secondInt = m[second]
		if secondInt == 0 {
			fmt.Print("Ошибка")
		} else {

			//Читаем операцию и записываем И переменную
			fmt.Print("Введите операцию: ")
			fmt.Fscan(os.Stdin, &operation)

			//Считаем результат и выводим в консоль
			switch operation {
			case "+":
				fmt.Print(firstInt + secondInt)
			case "-":
				fmt.Print(firstInt - secondInt)
			case "/":
				fmt.Print(firstInt / secondInt)
			case "*":
				fmt.Print(firstInt * secondInt)
			default:
				fmt.Print("Ошибка")
			}
		}
	}
}
