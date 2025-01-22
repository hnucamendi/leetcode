package main

import (
	"fmt"
	"math"
	"slices"
)

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
func sortedSquares(nums []int) []int {
	for i := range nums {
		if nums[i] < 0 {
			nums[i] = int(math.Abs(float64(nums[i])))
		}
		nums[i] = int(math.Pow(float64(nums[i]), 2))
	}
	slices.Sort(nums)
	return nums
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
