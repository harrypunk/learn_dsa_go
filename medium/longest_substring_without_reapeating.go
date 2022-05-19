//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package medium

func LengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		/*
		* traverse from right until a similar char is found
		* then every char in the substring is unique
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
