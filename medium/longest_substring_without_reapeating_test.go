package medium

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var s1 = "abcabcbb"
	var len1 = LengthOfLongestSubstring(s1)
	if len1 != 3 {
		t.Fail()
	}

	var s2 = ""
	var len2 = LengthOfLongestSubstring(s2)
	if len2 != 0 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstringHashMap(t *testing.T) {
	var s1 = "abcabcbb"
	var len1 = LengthOfLongestSubstringHash(s1)
	if len1 != 3 {
		t.Fail()
	}

}
