// https://leetcode-cn.com/problems/3sum

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right -= 1
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left += 1
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
