// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/

package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	imin, imax := 0, m
	var minRight, maxLeft int

	for imin <= imax {
		i := (imax + imin) / 2
		j := (m+n+1)/2 - i
		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums2[j] < nums1[i-1] {
			imax = i - 1
		} else {
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = int(math.Max(float64(nums1[i-1])/1.0, float64(nums2[j-1])))
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 5}, []int{2, 3, 4, 6}))
}
