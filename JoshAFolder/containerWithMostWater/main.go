// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.
//
// Example 1:
//
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
//
// Example 2:
//
// Input: height = [1,1]
// Output: 1
//
// Constraints:
//
//	n == height.length
//	2 <= n <= 105
//	0 <= height[i] <= 104
package main

import (
	"fmt"
	"math"
)

func calcWater(l, r, lv, rv int) int {
	min := math.Min(float64(lv), float64(rv))
	distance := float64(l - r)
	if distance < 0 {
		distance = math.Abs(float64(distance))
	}
	return int(min * distance)

}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxWater := 0
	for range height {
		water := 0
		if height[l] > height[r] {
			water = calcWater(l, r, height[l], height[r])
			r--
		} else if height[r] > height[l] {
			water = calcWater(l, r, height[l], height[r])
			l++
		} else {
			water = calcWater(l, r, height[l], height[r])
			if l+1 >= len(height) {
				if r-1 <= len(height) {
					break
				} else {
					r--
				}
			} else {
				l++
			}
		}

		if water > maxWater {
			maxWater = water
		}
	}
	return maxWater
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}
