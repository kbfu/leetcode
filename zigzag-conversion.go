// https://leetcode-cn.com/problems/zigzag-conversion

package main

import "fmt"

func convert(s string, numRows int) string {
	res := make([][]byte, numRows)
	flag := -1
	reversed := 0
	joined := ""
	for i := range s {
		if reversed == 0 {
			flag++
		} else {
			if flag > 0 {
				flag--
			}
		}
		if flag == numRows-1 {
			reversed = 1
		} else if flag == 0 {
			reversed = 0
		}
		res[flag] = append(res[flag], s[i])
	}

	for i := range res {
		joined = joined + string(res[i])
	}
	return joined
}

func main() {
	fmt.Println(convert("AB", 1))
}
