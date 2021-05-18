// https://leetcode-cn.com/problems/remove-element/
package main

import "fmt"

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func main() {
	fmt.Println(removeElement([]int{3, 3, 2, 3}, 3))
}
