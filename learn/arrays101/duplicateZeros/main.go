// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
package main

import "fmt"

func duplicateZeros(arr []int) {
	for i := range arr {
		if arr[i] == 0 {
      arr[i + 1] = arr 
		}
	}

	// fmt.Println(arr)
}

func main() {
	// Expected:
	// Output: [1,0,0,2,3,0,0,4]
	input := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(input)
}
