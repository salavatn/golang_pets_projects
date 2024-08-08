package main

import (
	"fmt" // Для ввода-вывода
	"os"
	"strconv" // Для преобразования чисел в строки
)

func main() {
	var userValue_A, userValue_B, userMath_Op string

	fmt.Printf("Введите пример (ex.: 2 + 5):  ")
	_, errorMsg := fmt.Scanln(&userValue_A, &userMath_Op, &userValue_B)

	if errorMsg != nil {
		fmt.Println("\nERROR: Некорректный ввод данных.\n")
		return
	}

	number_A, number_B := converter(userValue_A, userValue_B)
	math_Op := validator(userMath_Op)

	fmt.Printf("\n%d %s %d. Finish", number_A, math_Op, number_B)

	if math_Op == "+" {
		fmt.Printf("\nСумма чисел %d и %d = %d", number_A, number_B, number_A+number_B)
	} else if math_Op == "-" {
		fmt.Printf("\nВычитание чисел %d и %d = %d", number_A, number_B, number_A-number_B)
	} else if math_Op == "*" {
		fmt.Printf("\nУмножение чисел %d и %d = %d", number_A, number_B, number_A*number_B)
	} else if math_Op == "/" {
		if number_B != 0 {
			fmt.Printf("\nДеление чисел %d и %d = %d", number_A, number_B, number_A/number_B)
		} else {
			fmt.Println("\nERROR: Деление на ноль запрещено!")
		}
	}
}

func converter(a string, b string) (int, int) {
	num_a, errorMsg1 := strconv.Atoi(a)
	num_b, errorMsg2 := strconv.Atoi(b)

	if errorMsg1 != nil || errorMsg2 != nil {
		fmt.Printf("\nERROR: ошибка преобразования строки в числа\n")
		os.Exit(1)
	}
	return num_a, num_b
}

func validator(math_op string) string {
	switch math_op {
	case "+", "-", "*", "/":
		return math_op
	default:
		fmt.Println("\nERROR: 	Некорректный оператор")
		fmt.Println("\nMESSAGE: Используйте один из перечисленных \"+\", \"-\", \"*\", \"/\"\n")
		return ""
	}
}

// Функция конвертирования римских чисел в арабские
func roman_numbers(romanNumber string) int {

	var (
		arabicNumber int
		checking     bool
	)

	romanListNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15,
		"XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
	}

	// Конвертирование и возврат число
	if arabicNumber, checking = romanListNumerals[romanNumber]; checking {
		return arabicNumber
	} else {
		fmt.Println("Ошибка: некорректный ввод")
		return -1
	}
}
