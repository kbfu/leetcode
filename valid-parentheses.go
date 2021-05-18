// https://leetcode-cn.com/problems/valid-parentheses
package main

import "fmt"

func isValid(s string) bool {
	size := len(s)
	if size%2 != 0 {
		return false
	}
	var stack []byte

	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for i := range s {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[]{}()"))
	fmt.Println(isValid("[[{{}}]"))
}
