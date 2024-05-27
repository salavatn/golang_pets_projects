// Simple console app - calculator

package main

// Импорт модулей / пакетов
import (
	"fmt"     // Для ввода-вывода
	"strconv" // Для преобразования чисел в строки
)

// Точка входа в программу
func main() {
	fmt.Println(div(42, -1))
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
