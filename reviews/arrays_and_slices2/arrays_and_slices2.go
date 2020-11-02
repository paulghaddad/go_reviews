package arrays_and_slices2

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	tails := make([]int, len(slices))

	for i := 0; i < len(slices); i++ {
		tails[i] = Sum(slices[i][1:])
	}

	return tails
}
