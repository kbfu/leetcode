// https://leetcode-cn.com/problems/longest-palindromic-substring/description/

package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	palindrome := ""
	if len(s) <= 1 {
		return s
	}
	for i := range s {
		str1 := find(s, i, i)
		str2 := find(s, i, i+1)
		if len(str1) > len(palindrome) {
			palindrome = str1
		}
		if len(str2) > len(palindrome) {
			palindrome = str2
		}
	}
	return palindrome
}

func find(s string, left int, right int) string {
	if right < len(s) && s[left] == s[right] {
		for left > 0 && right < len(s)-1 {
			if s[left-1] == s[right+1] {
				left--
				right++
			} else {
				break
			}
		}
		return s[left : right+1]
	}
	return ""
}

func main() {
	fmt.Println(longestPalindrome("aaabaaaa"))
}
