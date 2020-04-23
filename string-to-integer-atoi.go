// https://leetcode-cn.com/problems/string-to-integer-atoi/description/

package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return 0
	}

	result := 0
	isNegative := false

	// check negative & validate string
	if !strings.Contains("1234567890-+", string(str[0])) {
		return 0
	}
	if str[0] == '-' {
		str = str[1:]
		isNegative = true
	} else if str[0] == '+' {
		str = str[1:]
	}

	// check alphabet
	for i := range str {
		if str[i]-'0' < 0 || str[i]-'0' > 9 {
			str = str[:i]
			break
		}
	}

	for i := 0; i < len(str); i++ {
		if int(str[i]-'0') > 9 || int(str[i]-'0') < 0 {
			break
		}
		if isNegative {
			if float64(int(str[i]-'0'))*math.Pow10(len(str)-i-1) > float64(math.MinInt32)*-1 {
				return math.MinInt32
			}
		} else {
			if float64(int(str[i]-'0'))*math.Pow10(len(str)-i-1) > float64(math.MaxInt32) {
				return math.MaxInt32
			}
		}
		result += int(str[i]-'0') * int(math.Pow10(len(str)-i-1))
	}

	if isNegative {
		result *= -1
	}
	if result < math.MinInt32 {
		return math.MinInt32
	} else if result > math.MaxInt32 {
		return math.MaxInt32
	}

	return result
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
