// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
// Example 1:
//
// Input: s = "abc", t = "ahbgdc"
// Output: true
//
// Example 2:
//
// Input: s = "axc", t = "ahbgdc"
// Output: false
//
// Constraints:
//
//	0 <= s.length <= 100
//	0 <= t.length <= 104
//	s and t consist only of lowercase English letters.
//
// Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?
package main

import (
	"fmt"
	"slices"
	"strings"
)

func isSubsequence(s string, t string) bool {
	sArr := strings.Split(s, "")
	for i := range len(t) {
		if slices.Contains(sArr, string(t[i])) {
			if i+1 <= len(sArr) {
				fmt.Println(sArr[i])
				sArr = append(sArr[:i], sArr[i+1])
			}
			sArr = []string{}
		}
	}
	if len(sArr) <= 0 {
		return true
	}
	return false

}

func main() {

}
