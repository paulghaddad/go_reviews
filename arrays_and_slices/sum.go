package arraysandslices

// Sum adds up the numbers in a slice
func Sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}

	return total
}

// SumAll returns a slice of the sums of each slice passed in
func SumAll(numsToSum ...[]int) []int {
	var sums []int

	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

// SumAllTails returns the sum of all elements except the first item in each
// input slice
func SumAllTails(nums ...[]int) []int {
	tailSums := make([]int, 0)
	for _, collection := range nums {
		var sum int
		if len(collection) == 0 {
			sum = 0
		} else {
			sum = Sum(collection[1:])
		}

		tailSums = append(tailSums, sum)
	}

	return tailSums
}
