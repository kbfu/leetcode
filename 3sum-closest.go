// https://leetcode-cn.com/problems/3sum-closest/
package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := nums[0] + nums[1] + nums[2]
	for i := range nums {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-res)) > math.Abs(float64(target-sum)) {
				res = sum
			}
			if sum > target {
				right -= 1
			} else if sum < target {
				left += 1
			} else {
				return res
			}
		}
	}
	return res
}
