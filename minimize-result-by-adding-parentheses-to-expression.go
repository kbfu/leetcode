// https://leetcode.cn/problems/minimize-result-by-adding-parentheses-to-expression/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minimizeResult(expression string) string {
	min := 0
	resultStr := ""
	exps := strings.Split(expression, "+")
	left, right := exps[0], exps[1]
	for i := len(left) - 1; i >= 0; i-- {
		for j := 1; j <= len(right); j++ {
			leftInsideParentheses, err := strconv.Atoi(left[i:])
			leftOutsideParentheses, err := strconv.Atoi(left[0:i])
			if err != nil {
				leftOutsideParentheses = 1
			}
			rightInsideParentheses, _ := strconv.Atoi(right[0:j])
			rightOutsideParentheses, err := strconv.Atoi(right[j:])
			if err != nil {
				rightOutsideParentheses = 1
			}
			result := leftOutsideParentheses * (leftInsideParentheses + rightInsideParentheses) * rightOutsideParentheses
			if min == 0 || min > result {
				min = result
				resultStr = fmt.Sprintf("%s(%s+%s)%s", left[0:i], left[i:], right[0:j], right[j:])
			}
		}
	}
	return resultStr
}

func main() {
	print(minimizeResult("22+1"))
}
