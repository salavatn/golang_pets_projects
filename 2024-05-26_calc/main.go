package main

// Импорт модулей / пакетов
import (
	"fmt"     // Для ввода-вывода
	"strconv" // Для преобразования чисел в строки
)

func main() {
	var strValueA, strValueB, strMathOperator string
	fmt.Print("Ваш пример: ")
	_, errorMsg := fmt.Scanf("%s %s %s", &strValueA, &strMathOperator, &strValueB)

	if errorMsg != nil {
		fmt.Println("\tОшибка: некорректный ввод")
		return
	}

	valueA, errorMsg1 := strconv.Atoi(strValueA)
	valueB, errorMsg2 := strconv.Atoi(strValueB)

	if errorMsg1 != nil || errorMsg2 != nil {
		fmt.Println("\tОшибка: некорректный ввод чисел")
		return
	}

	if strMathOperator == "+" {
		fmt.Printf("Ваш результат: %d\n\n", add(valueA, valueB))
	} else if strMathOperator == "-" {
		fmt.Printf("Ваш результат: %d\n\n", sub(valueA, valueB))
	} else if strMathOperator == "*" {
		fmt.Printf("Ваш результат: %d\n\n", mul(valueA, valueB))
	} else if strMathOperator == "/" {
		fmt.Printf("Ваш результат: %s\n\n", div(valueA, valueB))
	} else {
		fmt.Println("\tОшибка: некорректная операция")
		return
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
	result_str := strconv.FormatFloat(float64(result_flt), 'f', 6, 64) // Преобразование числа в строку
	return string(result_str)
}
