// https://leetcode-cn.com/problems/longest-common-prefix/

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) < 2 {
		return strs[0]
	}
	shortest := len(strs[0])
	index := 0
	for _, s := range strs {
		if shortest > len(s) {
			shortest = len(s)
		}
	}

	for {
		for i := 1; i < len(strs); i++ {
			if index == shortest || strs[0][index] != strs[i][index] {
				return strs[0][:index]
			}
		}
		index += 1
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"a"}))
}
