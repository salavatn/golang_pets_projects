package main

import (
	"fmt" // Для ввода-вывода
	"os"
	"strconv"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15,
	"XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
	"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25,
	"XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
	"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35,
	"XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
	"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45,
	"XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
	"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55,
	"LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
	"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65,
	"LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
	"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75,
	"LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
	"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85,
	"LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95,
	"XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
}
var userValueA, userValueOp, userValueB string
var data = map[string]interface{}{
	"number_a": 0,     // int
	"number_b": 0,     // int
	"type_a":   false, // Roman (true) or Araibc (false)
	"type_b":   false, // Roman (true) or Araibc (false)
	"math_op":  "+",   // + - / *
	"result":   0,     // int
}

func main() {
	fmt.Printf("Input:\t")
	_, userInputError := fmt.Scanln(&userValueA, &userValueOp, &userValueB)
	if userInputError != nil {
		fmt.Printf("Output: Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	check_roman(userValueA, userValueB)
	check_math_op(userValueOp)

	roman_a := data["type_a"].(bool)
	roman_b := data["type_b"].(bool)

	if roman_a == false && roman_b == false {
		check_arabic(userValueA, userValueB)
	} else if roman_a == true && roman_b == false {
		fmt.Println("Output: Выдача паники, так как используются одновременно разные системы счисления.")
		return
	} else if roman_a == false && roman_b == true {
		fmt.Println("Output: Выдача паники, так как используются одновременно разные системы счисления.")
		return
	}

	math(data["number_a"].(int), data["math_op"].(string), data["number_b"].(int))

	result := data["result"].(int)

	if roman_a == true && roman_b == true {
		for roman, arabic := range romanNumerals {
			if arabic == result {
				fmt.Printf("Output:\t%s\n\n", roman)
			}
		}
	} else {
		fmt.Printf("Output:\t%d\n\n", result)
	}

}

func check_roman(a string, b string) {
	if value, check := romanNumerals[a]; check {
		data["number_a"] = value
		data["type_a"] = check
	} else {
		data["number_a"] = 0
		data["type_a"] = false
	}

	if value, check := romanNumerals[b]; check {
		data["number_b"] = value
		data["type_b"] = check
	} else {
		data["number_b"] = 0
		data["type_b"] = false
	}
}

func check_arabic(a string, b string) {
	num_a, convertingErr1 := strconv.Atoi(a)
	num_b, convertingErr2 := strconv.Atoi(b)

	if convertingErr1 != nil || convertingErr2 != nil {
		fmt.Printf("\nOutput: Выдача паники, так как строка не является математической операцией.")
		os.Exit(0)
	}
	if num_a >= 1 && num_a <= 10 && num_b >= 1 && num_b <= 10 {
		data["number_a"] = num_a
		data["type_a"] = false
		data["number_b"] = num_b
		data["type_b"] = false
	} else {
		panic("Output: Использовать только числа 1,2, ..., 10!")
	}
}

func check_math_op(operation string) {
	switch operation {
	case "+", "-", "*", "/":
		// fmt.Printf("\n\nOperator is: %s\n\n", operation)
		data["math_op"] = operation
	default:
		fmt.Println("Output: Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}
}

func math(a int, op string, b int) {
	roman_a := data["type_a"].(bool)
	roman_b := data["type_b"].(bool)

	if op == "+" {
		data["result"] = a + b
	} else if op == "-" {
		if roman_a == true && roman_b == true {
			fmt.Println("Output: Выдача паники, так как в римской системе нет отрицательных чисел.")
			os.Exit(0)
		}
		data["result"] = a - b
	} else if op == "*" {
		data["result"] = a * b
	} else if op == "/" {
		if b == 0 {
			fmt.Println("Output: Выдача паники, так как делить на 0 запрещено!")
			os.Exit(0)
		} else if roman_a == true && roman_b == true && (a/b < 1) {
			panic("Паника, так как результат меньше единицы в римской системе!")
		}
		data["result"] = a / b
	}
}
