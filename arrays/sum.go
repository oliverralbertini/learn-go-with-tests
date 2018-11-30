package arrays

func Sum(nums []int) (sum int) {
	for _, val := range nums {
		sum += val
	}
	return
}

func SumAll(numSlcs ...[]int) (sums []int) {
	for _, slc := range numSlcs {
		sums = append(sums, Sum(slc))
	}
	return
}

func SumAllTails(numSlcs ...[]int) (sums []int) {
	for _, slc := range numSlcs {
		tail := slc[1:]
		sums = append(sums, Sum(tail))
	}
	return
}

func SumAllHeads(numSlcs ...[]int) (sums []int) {
	for _, slc := range numSlcs {
		head := slc[0 : len(slc)-1]
		sums = append(sums, Sum(head))
	}
	return
}

func SumAllIns(numSlcs ...[]int) (sums []int) {
	for _, slc := range numSlcs {
		head := slc[1 : len(slc)-1]
		sums = append(sums, Sum(head))
	}
	return
}
