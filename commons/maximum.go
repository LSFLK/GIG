package commons

/**
return maximum of a positive number slice
 */
func Maximum(list []int) int {
	max := -1
	for _, num := range list {
		if num > max || max == -1 {
			max = num
		}
	}
	return max
}