// https://leetcode.com/problems/move-zeroes/description/
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
// Example 1:
//
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example 2:
//
// Input: nums = [0]
// Output: [0]
//
// Constraints:
//
//	1 <= nums.length <= 104
//	-231 <= nums[i] <= 231 - 1
//
// Follow up: Could you minimize the total number of operations done?
package main

import "fmt"

func moveZeroes(nums []int) {
	i := 0
	j := i + 1

	for range len(nums) {
		if j >= len(nums) {
			break
		}

		if i >= len(nums) {
			break
		}

		if nums[j] == 0 {
			if nums[i] != 0 {
        t := nums[i]
				nums[i] = nums[j]
				nums[j] = t
			}
		}

		if nums[j] != 0 {
			if nums[i] == 0 {
				nums[i] = nums[j]
				nums[j] = 0
			}
			i++
		}
		j++
	}
	fmt.Println(nums)
}

func main() {
	input1 := []int{0, 1, 0, 3, 12}
	input2 := []int{2, 1}
	input3 := []int{1, 0, 1}
	input4 := []int{1, 2, 3, 1}
	input5 := []int{1, 0}
	input6 := []int{1, 0, 0}
	moveZeroes(input1)
	moveZeroes(input2)
	moveZeroes(input3)
	moveZeroes(input4)
	moveZeroes(input5)
	moveZeroes(input6)
}
