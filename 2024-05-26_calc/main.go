package main

// Импорт модулей / пакетов
import (
	"fmt"     // Для ввода-вывода
	"strconv" // Для преобразования чисел в строки
)

func main() {
	// MAIN. Part 1: Объявление переменных
	var userValue_A, userValue_B, userMath_Op string

	// MAIN. Part 2: Ввод примера, есть ли ошибка и сохрание в переменные
	fmt.Print("Ваш пример: ")
	_, errorMsg := fmt.Scanf("%s %s %s", &userValue_A, &userMath_Op, &userValue_B)

	if errorMsg != nil {
		fmt.Println("\tОшибка: некорректный ввод")
		return
	}

	// MAIN. Part 3: Объявление переменных для чисел и конвертация строк в числа
	arabicNumA, errorMsg1 := strconv.Atoi(userValue_A)
	arabicNumB, errorMsg2 := strconv.Atoi(userValue_B)

	// // MAIN. Part 4: Если введены римские числа, то конвертировать их в арабские
	// if errorMsg1 != nil {
	// 	arabicNumA = roman_numbers(userValue_A)
	// }

	// MAIN. Part 4: Если конвертация успешна, то выполнить операцию
	if errorMsg1 == nil && errorMsg2 == nil {
		if userMath_Op == "+" {
			fmt.Printf("Ваш результат: %d\n\n", add(arabicNumA, arabicNumB))
		} else if userMath_Op == "-" {
			fmt.Printf("Ваш результат: %d\n\n", sub(arabicNumA, arabicNumB))
		} else if userMath_Op == "*" {
			fmt.Printf("Ваш результат: %d\n\n", mul(arabicNumA, arabicNumB))
		} else if userMath_Op == "/" {
			fmt.Printf("Ваш результат: %s\n\n", div(arabicNumA, arabicNumB))
		} else {
			fmt.Println("\tОшибка: некорректная операция\n")
			return
		}
	} else if (errorMsg1 == nil && errorMsg2 != nil) || (errorMsg1 != nil && errorMsg2 == nil) {
		fmt.Println("\tОшибка: некорректный пример")
		return
	}

	// MAIN. Part 5: Объявление переменных для римских чисел и конвертация string в int
	var romanNumA, romanNumB int
	romanNumA = roman_numbers(userValue_A)
	romanNumB = roman_numbers(userValue_B)

	// MAIN. Part 6: Если конвертация успешна, то выполнить операцию
	if romanNumA != -1 && romanNumB != -1 {
		if userMath_Op == "+" {
			fmt.Printf("Ваш результат: %d\n\n", add(romanNumA, romanNumB))
		} else if userMath_Op == "-" {
			fmt.Printf("Ваш результат: %d\n\n", sub(romanNumA, romanNumB))
		} else if userMath_Op == "*" {
			fmt.Printf("Ваш результат: %d\n\n", mul(romanNumA, romanNumB))
		} else if userMath_Op == "//" {
			fmt.Printf("Ваш результат: %s\n\n", div(romanNumA, romanNumB))
		} else {
			fmt.Println("\tОшибка: некорректная операция\n")
			return
		}
	}

}

// Функции сложения (add), вычитания (sub), умножения (mul) и деления (div) двух чисел.
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) string {
	if b == 0 {
		return "Ошибка: деление на ноль запрещено!"
	}
	result_flt := float64(a) / float64(b)
	// result_flt := a % b // Деление без остатка
	result_str := strconv.FormatFloat(float64(result_flt), 'f', 6, 64) // Преобразование числа в строку
	return string(result_str)
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
