// https://leetcode-cn.com/problems/container-with-most-water
package main

import "math"

func maxArea(height []int) int {
	left, right := 0.0, float64(len(height)-1)
	maxArea := 0.0
	for left < right {
		leftHeight := height[int(left)]
		rightHeight := height[int(right)]

		area := math.Min(float64(leftHeight), float64(rightHeight)) * (right - left)
		maxArea = math.Max(maxArea, area)
		if leftHeight > rightHeight {
			right -= 1
		} else if leftHeight < rightHeight {
			left += 1
		} else {
			left += 1
		}
	}
	return int(maxArea)
}
