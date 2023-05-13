package variativevariable

func MergeNumberLists(numberLists ...[]int) []int{
	my_slice := make([]int,0) 
	for _, slice := range numberLists {
		for _, val := range slice {
			my_slice = append(my_slice, val)
		}
	}
	return my_slice
}
