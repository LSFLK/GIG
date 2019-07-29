package utility

/**
return minimum of a positive number slice
 */
func Minimum(list []int) int {
	min := -1
	for _, num := range list {
		if num < min || min == -1 {
			min = num
		}
	}
	return min
}
