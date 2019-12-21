// https://leetcode-cn.com/problems/roman-to-integer/
package main

import "fmt"

var (
	mapping = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	specialMapping = map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
)

func romanToInt(s string) int {
	result := 0
	pass := false
	for i := range s {
		if pass == false {
			if i+2 <= len(s) {
				v, ok := specialMapping[s[i:i+2]]
				if ok {
					result += v
					pass = true
				} else {
					result += mapping[s[i:i+1]]
				}
			} else if i+1 <= len(s) {
				result += mapping[s[i:i+1]]
			}
		} else {
			pass = false
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
}
