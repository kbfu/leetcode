// https://leetcode-cn.com/problems/string-to-integer-atoi/description/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	numStr := ""
	if len(str) == 0 {
		return 0
	}
	if str[0] == 45 || str[0] == 43 {
		numStr += string(str[0])
		for i := 1; i < len(str); i++ {
			if str[i] <= 57 && str[i] >= 48 {
				numStr += string(str[i])
			} else {
				break
			}
		}
	} else if str[0] <= 57 && str[0] >= 48 {
		for i := 0; i < len(str); i++ {
			if str[i] <= 57 && str[i] >= 48 {
				numStr += string(str[i])
			} else {
				break
			}
		}
	} else {
		return 0
	}
	num, _ := strconv.Atoi(numStr)
	if num > math.MaxInt32 {
		return math.MaxInt32
	} else if num < math.MinInt32 {
		return math.MinInt32
	} else {
		return num
	}
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
