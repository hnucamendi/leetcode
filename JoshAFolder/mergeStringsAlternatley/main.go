// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
//
// Return the merged string.
//
// Example 1:
//
// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r
//
// Example 2:
//
// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b
// word2:    p   q   r   s
// merged: a p b q   r   s
//
// Example 3:
//
// Input: word1 = "abcd", word2 = "pq"
// Output: "apbqcd"
// Explanation: Notice that as word1 is longer, "cd" is appended to the end.
// word1:  a   b   c   d
// word2:    p   q
// merged: a p b q c   d
//
// Constraints:
//
//	1 <= word1.length, word2.length <= 100
//	word1 and word2 consist of lowercase English letters.
package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1, word2 string) string {
	i, j := 0, 0
  res := make([]byte, len(word1)+len(word2))
	for i < len(word1) && j < len(word2) {
		res = append(res, word1[i])
		res = append(res, word2[j])
		i++
		j++
	}


	res = append(res, word1[i:]...)
	res = append(res, word2[j:]...)

  sRes := strings.Trim(strings.Join(strings.Split(string(res),""), ""),"\u0000")
  return sRes
}

func main() {
	inputs := []struct {
		word1 string
		word2 string
	}{
		{

			word1: "abc",
			word2: "pqr",
		},
		{
			word1: "ab",
			word2: "pqrs",
		},
		{
			word1: "abcd",
			word2: "pq",
		},
	}

	for _, k := range inputs {
		fmt.Println(mergeAlternately(k.word1, k.word2))
	}
}

// ### DNF 2hrs ###
// 	var exW []byte
// 	if len(w1)-len(w2) != 0 {
// 		if len(w1)-len(w2) < 0 {
// 			tword1 := []byte(w1)
// 			tword2 := []byte(w2)
// 			ex := len(tword2) - len(tword1)
//
// 			exW = tword2[ex:]
// 			mergeStr = append(mergeStr, tword1...)
// 			mergeStr = append(mergeStr, tword2[:ex]...)
// 		} else {
// 			tword1 := []byte(w1)
// 			tword2 := []byte(w2)
// 			ex := len(tword1) - len(tword2)
//
// 			exW = tword1[ex:]
// 			mergeStr = append(mergeStr, tword1[:ex]...)
// 			mergeStr = append(mergeStr, tword2...)
// 		}
// 	} else {
// 		mergeStr = append(mergeStr, w1...)
// 		mergeStr = append(mergeStr, w2...)
// 	}
// 	return mergeStr, exW
// }
//
// func mergeAlternately(word1 string, word2 string) string {
// 	var mergeStr, exW = equalizeLength(word1, word2)
// 	fmt.Println(string(mergeStr))
//
// 	str1 := []byte{}
//   var sn int
//   sn = len(word1)
// 	if len(word1) > len(word2) {
// 		sn = len(word1) - len(exW)
// 	}
// 	var j = sn
// 	var i = 0
// 	var mergeI = false
// 	for range len(mergeStr) {
// 		fmt.Println(i, j)
// 		fmt.Println(string(mergeStr[i]), string(mergeStr[j]))
// 		mergeI = !mergeI
// 		if mergeI {
// 			str1 = append(str1, mergeStr[i])
// 			if i+1 > sn {
// 				continue
// 			}
// 			i++
// 		}
//
// 		if !mergeI {
// 			str1 = append(str1, mergeStr[j])
// 			if j+1 > len(mergeStr) {
// 				break
// 			}
// 			j++
// 		}
// 	}
// 	if len(exW) > 0 {
// 		str1 = append(str1, exW...)
// 	}
// 	return string(str1)
// }
//
// func main() {
// 	inputs := []struct {
// 		word1 string
// 		word2 string
// 	}{
// 		// {
// 		//
// 		// 	word1: "abc",
// 		// 	word2: "pqr",
// 		// },
// 		// {
// 		// 	word1: "ab",
// 		// 	word2: "pqrs",
// 		// },
// 		{
// 			word1: "abcd",
// 			word2: "pq",
// 		},
// 	}
//
// 	for _, k := range inputs {
// 		fmt.Println(mergeAlternately(k.word1, k.word2))
// 	}
// }
