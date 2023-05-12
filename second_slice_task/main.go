package secondslicetask

func IntsCopy(src []int, maxLen int) []int {
	if maxLen < 0 {
		return []int{}
	}
	if maxLen > len(src) {
		maxLen = len(src)
	} 
	copy_slice := make([]int, maxLen)
	copy(copy_slice, src)
	return copy_slice
}