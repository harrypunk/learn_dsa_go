//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package medium

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

	/*
	 * left is the index that last time the same char occured
	 *  then length = right - left,
	 * Edge case: index 0, left has to start with -1
	 */
	left := -1
	maxLen := 1
	// char : last occur indices
	var charMap = map[byte]int{
		s[0]: 0,
	}

	for right := 1; right < len(s); right++ {
		var rightChar = s[right]
		if lastOccurIndex, ok := charMap[rightChar]; ok {
			if lastOccurIndex > left {
				left = lastOccurIndex
			}
		}
		var len1 = right - left
		if len1 > maxLen {
			maxLen = len1
		}
		charMap[rightChar] = right

	}

	return maxLen
}
