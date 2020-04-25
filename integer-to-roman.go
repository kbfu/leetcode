package main

import "fmt"

func intToRoman(num int) string {
	romanInt := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanStr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	index := 0

	result := ""
	for {
		if num-romanInt[index] >= 0 {
			result += romanStr[index]
			num = num - romanInt[index]
		} else {
			index += 1
			if index == len(romanStr) {
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(intToRoman(6))
}
