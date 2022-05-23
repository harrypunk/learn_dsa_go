package easy

func ReverseString(s []byte) {
	var slen = len(s)

	var t byte
	for left, right := 0, slen-1; left < right; {
		t = s[left]
		s[left] = s[right]
		s[right] = t
		left++
		right--
	}

}
