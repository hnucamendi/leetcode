package main

import (
	"fmt"
	"slices"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	answer := make([][]int, 2)

	fmt.Println(nums1)
	fmt.Println(nums2)

	for i := range nums2 {
		if nums1[i] == len(nums1) {
			break
		}
		if !slices.Contains(nums2, nums1[i]) {
			if slices.Contains(answer[0], nums1[i]) {
				continue
			}
			answer[0] = append(answer[0], nums1[i])
		}
	}

	for i := range nums1 {
		if nums2[i] == len(nums2) {
			break
		}
		if !slices.Contains(nums1, nums2[i]) {
			if slices.Contains(answer[1], nums2[i]) {
				continue
			}
			answer[1] = append(answer[1], nums2[i])
		}
	}

	return answer
}

// Output: [[3],[]]
func main() {
	nums1 := []int{-68, -80, -19, -94, 82, 21, -43}
	nums2 := []int{-63}
	fmt.Println(findDifference(nums1, nums2))
}
