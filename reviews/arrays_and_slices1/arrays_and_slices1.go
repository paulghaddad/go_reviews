package arrays_and_slices1

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumSlices(slices ...[]int) []int {
	numSlices := len(slices)
	sums := make([]int, numSlices)

	for i, slice := range slices {
		sums[i] = Sum(slice)
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	tailSums := make([]int, 0)

	for _, slice := range slices {
		if len(slice) >= 1 {
			tailSums = append(tailSums, Sum(slice[1:]))
		}
	}

	return tailSums
}
