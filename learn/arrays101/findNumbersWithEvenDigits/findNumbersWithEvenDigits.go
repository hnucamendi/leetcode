package main

import (
	"fmt"
)

// Given an array nums of integers, return how many of them contain an even number of digits.
func calcDigitSize(c, n int) int {
	if n == 0 {
		return c
	}
	c++
	return calcDigitSize(c, n/10)
}

func findNumbers(nums []int) int {
	count := 0
	for _, k := range nums {
		if calcDigitSize(0, k)&1 == 0 {
			count++
		}
	}

	return count
}

func main() {
	input := []int{555, 901, 482, 1771}
	fmt.Println(findNumbers(input))

}
