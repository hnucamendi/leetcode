// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
//
// Example 1:
//
// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
//
// Example 2:
//
// Input: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.
//
// Example 3:
//
// Input: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
//
// Constraints:
//
//	   1 <= nums.length <= 5 * 105
//	   -231 <= nums[i] <= 231 - 1
//
//	Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?
package main

import "fmt"

func increasingTriplet(nums []int) bool {
	i := 0
	j := i + 1
	k := j + 1

  var ti, tj, tk int

	for range nums {
		if k > len(nums)-1 {
			return false
		}

    if nums[i] >= nums[j] {
      i++
    }else {
      ti = nums[i]
    }

    if nums[j] <= nums[i] || nums[j] >= nums[k] {
      j++
    }else {
      tj = nums[j]
    }

    if nums[k] <= nums[j] {
      k++
    }else {
      tk = nums[k]
    }

    if ti < tj && tk > tj {
      return true
    }

	}
	return false
}

func main() {
	// input1 := []int{1, 2, 3, 4, 5}
	// input2 := []int{5, 4, 3, 2, 1}
	// input3 := []int{2, 1, 5, 0, 4, 6}
	input4 := []int{20, 100, 10, 12, 5, 13}
	// 20, 100, 10
	// 100, 10, 12
	// 10,12,5
	// 12,5,13
	// fmt.Println(increasingTriplet(input1))
	// fmt.Println(increasingTriplet(input2))
	// fmt.Println(increasingTriplet(input3))
	fmt.Println(increasingTriplet(input4))
}
