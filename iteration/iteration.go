package iteration

func Repeat(repeatCount int, toRepeat string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += toRepeat
	}
	return
}

func Reverse(toReverse string) string {
	reversed := []rune(toReverse)
	for i, j := 0, len(toReverse)-1; j > i; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}
