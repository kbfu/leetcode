// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
package main

import "fmt"

var phone = map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}

func deep(current string, digits string, res *[]string) {
	if len(digits) == 0 {
		return
	}
	for j := range phone[string(digits[0])] {
		if len(digits)-1 == 0 {
			*res = append(*res, current+string(phone[string(digits[0])][j]))
		}
		deep(current+string(phone[string(digits[0])][j]), digits[1:], res)
	}
}

func letterCombinations(digits string) []string {
	var res []string
	deep("", digits, &res)
	return res
}

func main() {
	fmt.Println(letterCombinations("233"))
}
