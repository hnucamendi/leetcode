package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(U int, weight []int) int {
	// Implement your solution here
	var i, j int
	carCounter := 0
	for range weight {
		j = i + 1
		if j > len(weight)-1 {
			break
		}

		// fmt.Println(i, j, "TAMO")

		carsOnBridge := weight[i] + weight[j]
		fmt.Printf("%v\t %v\t %v\n", weight[i], weight[j], carsOnBridge)
		if carsOnBridge > U {
			if weight[i] > weight[j] {
				weight = append(weight[:i], weight[i+1:]...)
			} else if weight[i] < weight[j] {
				weight = append(weight[:j], weight[j+1:]...)
			} else {
				weight = append(weight[:i], weight[i+1:]...)
			}
		} else {
			carCounter += 1
			i++
		}
		// fmt.Printf("%v\t %v\t %v\t %v\n", weight[i], weight[j], carsOnBridge, carCounter)
		// fmt.Println(carCounter)
	}

	fmt.Println(weight)
	return len(weight)
}

// func test() {

// }

func main() {
	U := 7
	weight := []int{7, 6, 5, 2, 7, 4, 5, 4} // 5
	fmt.Println(Solution(U, weight))
}
