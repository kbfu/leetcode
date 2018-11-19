// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/description/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(nums[0:removeDuplicates(nums)])
}

func removeDuplicates(nums []int) int {
	flag := 0
	if len(nums) == 0 {
		return 0
	}
	for i := range nums {
		if nums[i] != nums[flag] {
			flag++
			nums[flag] = nums[i]
		}
	}
	flag += 1
	return flag
}
