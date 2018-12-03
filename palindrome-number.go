// https://leetcode-cn.com/problems/palindrome-number/description/

package main

import "fmt"

func isPalindrome(x int) bool {
	reversed := 0
	if x < 0 {
		return false
	} else {
		temp := x
		for temp != 0 {
			pop := temp % 10
			temp /= 10
			reversed = reversed*10 + pop
		}
	}
	if x == reversed {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(121))
}
