// https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/22/

package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	flag := 0
	min := 0
	max := 0

	for i := range prices {
		if min == 0 || prices[i] < min {
			min = prices[i]
			flag = i + 1
		}
	}
	if flag == len(prices) {
		return 0
	}
	for _, v := range prices[flag:] {
		if max == 0 || v > max {
			max = v
		}
	}
	return max - min
}
