package arrays

func Sum(nums [5]int) (sum int) {
	for _, val := range nums {
		sum += val
	}
	return
}
