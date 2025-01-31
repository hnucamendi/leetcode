// You are given an integer array nums and an integer k.
//
// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
//
// Return the maximum number of operations you can perform on the array.
//
//
//
// Example 1:
//
// Input: nums = [1,2,3,4], k = 5
// Output: 2
// Explanation: Starting with nums = [1,2,3,4]:
// - Remove numbers 1 and 4, then nums = [2,3]
// - Remove numbers 2 and 3, then nums = []
// There are no more pairs that sum up to 5, hence a total of 2 operations.
//
// Example 2:
//
// Input: nums = [3,1,3,4,3], k = 6
// Output: 1
// Explanation: Starting with nums = [3,1,3,4,3]:
// - Remove the first two 3's, then nums = [1,4,3]
// There are no more pairs that sum up to 6, hence a total of 1 operation.
//
//
//
// Constraints:
//
//     1 <= nums.length <= 105
//     1 <= nums[i] <= 109
//     1 <= k <= 109

package main

import "fmt"

func maxOperations(nums []int, k int) int {
	i, j := 0, len(nums)-1
	count := 0
	for range nums {
    fmt.Println(i, j, nums)
		if (i > len(nums) || j < len(nums)) && nums[i]+nums[j] == k {
			count++
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums[j:], nums[:j-1]...)
			fmt.Println(nums)
		}
		i++
		j--
	}
	return count
}

func main() {
	input := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(input, k))

}
