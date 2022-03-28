package easy

func ContainsDuplicate(nums []int) bool {
	len1 := len(nums)
	dupMap := make(map[int]bool, len1)
	for i := 0; i < len1; i++ {
		key := nums[i]
		if dupMap[key] {
			return true
		} else {
			dupMap[key] = true
		}
	}
	return false
}
