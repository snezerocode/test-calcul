package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите математическое выражение (например, 1 + 5): ")
	scanner.Scan()

	parts := strings.Split(scanner.Text(), " ")

	if len(parts) != 3 {
		fmt.Println("Некорректное выражение. Используйте формат 'число оператор число' (например, 1 + 5).")
		return
	}

	left, isRoman1 := parseInput(parts[0])
	operator := parts[1]
	right, isRoman2 := parseInput(parts[2])

	if (isRoman1 && !isRoman2) || (!isRoman1 && isRoman2) {
		fmt.Println("Некорректные числа или операторы.")
		return
	}

	var result int

	switch operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right == 0 {
			fmt.Println("Деление на ноль недопустимо.")
			return
		}
		result = left / right
	default:
		fmt.Println("Неподдерживаемый оператор. Поддерживаются +, -, *, /.")
		return
	}

	if isRoman1 || isRoman2 {
		if result <= 0 {
			fmt.Println("Результат не может быть отрицательным для римских цифр.")
			return
		}
		romanResult := intToRoman(result)
		fmt.Printf("Результат: %s\n", romanResult)
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}

func parseInput(input string) (int, bool) {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if num, ok := romanNumerals[input]; ok {
		return num, true
	} else if num, err := strconv.Atoi(input); err == nil {
		return num, false
	} else {
		return 0, false
	}

}

func intToRoman(num int) string {
	vals := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	symbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	roman := ""

	for i := 8; i >= 0; {
		if num >= vals[i] {
			roman += symbols[i]
			num -= vals[i]
		} else {
			i--
		}
	}

	return roman
}
