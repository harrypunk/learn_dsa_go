//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package medium

import "fmt"

func LengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		/*
					* traverse from right until a similar char is found
					* then every char in the substring is unique
			    * O(n^2)
		*/
		for i := right - 1; i >= left; i-- {
			if s[i] == s[right] {
				left = i + 1
				break
			}
		}
		var len1 = right - left + 1
		if len1 > maxLen {
			maxLen = len1
		}
	}

	return maxLen
}

func LengthOfLongestSubstringHash(s string) int {
	var n = len(s)
	if n < 2 {
		return n
	}

	left := 0
	maxLen := 1
	// char : last occur index
	var charMap = map[byte]int{}

	for right := 0; right < len(s); right++ {
		var rightChar = s[right]
		var rightCharIndex = 0
		if _, ok := charMap[rightChar]; ok {
			rightCharIndex = charMap[rightChar]
		}
		if rightCharIndex > left {
			left = rightCharIndex
		}
		var len1 = right - left + 1
		if len1 > maxLen {
			maxLen = len1
		}
		charMap[rightChar] = right + 1

		for i, v := range charMap {
			fmt.Printf("%c:%d ", i, v)
		}
		fmt.Printf("sub: %v", s[left:right+1])
		fmt.Println()

	}

	return maxLen
}
