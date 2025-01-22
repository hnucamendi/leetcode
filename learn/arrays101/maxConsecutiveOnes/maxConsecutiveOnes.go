package maxConsecutiveOnes

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	largestCount := 0
	for _, v := range nums {
		if v^1 == 0 {
			count++
		} else {
			if count > largestCount {
				largestCount = count
			}
			count = 0
		}
	}

	if count > largestCount {
		largestCount = count
	}

	return largestCount
}

// func main() {
// 	input := []int{1, 0, 1, 1, 0, 1}
// 	//expect 2
// 	fmt.Println(findMaxConsecutiveOnes(input))
// }
