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
		sum := 0
		if len(slc) > 0 {
			tail := slc[1:]
			sum = Sum(tail)
		}
		sums = append(sums, sum)
	}
	return
}

func SumAllHeads(numSlcs ...[]int) (sums []int) {
	for _, slc := range numSlcs {
		sum := 0
		if len(slc) > 0 {
			head := slc[0 : len(slc)-1]
			sum = Sum(head)
		}
		sums = append(sums, sum)
	}
	return
}

func SumAllIns(numSlcs ...[]int) (sums []int) {
	for _, slc := range numSlcs {
		sum := 0
		if len(slc) > 2 {
			head := slc[1 : len(slc)-1]
			sum = Sum(head)
		}
		sums = append(sums, sum)
	}
	return
}
