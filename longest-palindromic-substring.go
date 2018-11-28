// https://leetcode-cn.com/problems/longest-palindromic-substring/description/

package main

import (
	"bytes"
	"fmt"
)

func longestPalindrome(s string) string {
	palindrome := ""
	if len(s) <= 1 {
		return s
	}
	for i := range s {
		lastIndex := bytes.LastIndex([]byte(s), []byte{s[i]})
		if lastIndex >= i {
			str := s[i : lastIndex+1]
			strLen := len(str)
			if strLen%2 == 0 {
				if str[0:strLen/2] == str[strLen/2:strLen] {
					if strLen > len(palindrome) {
						palindrome = str
					}
				}
			} else {
				midIndex := strLen / 2
				if midIndex > 0 {
					for i := 0; i < midIndex; i++ {
						leftIndex := i
						rightIndex := strLen - i - 1
						if str[leftIndex] != str[rightIndex] {
							break
						}
						if i == midIndex-1 && strLen > len(palindrome) {
							palindrome = str
						}
					}
				} else {
					if strLen > len(palindrome) {
						palindrome = str
					}
				}
			}
		}
	}
	return palindrome
}

func main() {
	fmt.Println(longestPalindrome("aaabaaaa"))
}
