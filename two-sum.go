// https://leetcode-cn.com/problems/two-sum/description/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		_, ok := m[nums[i]]
		if ok {
			return []int{m[nums[i]], i}
		}
		m[another] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
