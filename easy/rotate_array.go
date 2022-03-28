package easy

func Rotate(nums []int, k int) {
	var nlen = len(nums)
	arr := make([]int, nlen)
	copy(arr, nums)
	for i := 0; i < nlen; i++ {
		next := (i + k) % nlen
		nums[next] = arr[i]
	}
}
