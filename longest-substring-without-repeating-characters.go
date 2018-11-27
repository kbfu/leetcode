// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/

package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	max := 0
	start := 0
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		prev := string(s[start:i])
		char := string(s[i])
		if strings.Count(prev, char) > 0 {
			if len(prev) > max {
				max = len(prev)
			}
			start = m[char] + 1
		} else {
			if len(prev)+1 > max {
				max = len(prev) + 1
			}
		}
		m[char] = i
	}
	return max
}

func main() {
	println(lengthOfLongestSubstring(" "))
	println(lengthOfLongestSubstring("abcabc"))
	println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring(""))
	println(lengthOfLongestSubstring("au"))
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring("nfpdmpi"))
}
