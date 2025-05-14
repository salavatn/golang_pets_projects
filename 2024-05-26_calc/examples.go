package main

import (
	"fmt"
)

var romanNumerals1 = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15,
	"XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
}

func main() {

	romanNum("XVII")

}

func romanNum(userValue string) {

	// Проверяем, есть ли указанная дата в массиве
	if value, check := romanNumerals1[userValue]; check {
		fmt.Printf("Значение для %s: %d\n", userValue, value)
		fmt.Printf("\nValue =  %s\nCheck = %s\n", value, check)

	} else {
		fmt.Println("Ошибка: указанного римского числа нет в массиве")
	}
}
