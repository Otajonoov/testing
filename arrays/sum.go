package arrays


// For array
func Sum(numbers [5]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// For Slice
func Sum2(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, number := range numbers {
		sums = append(sums, Sum2(number))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum2(tail))
		}
	}
	return sums
}
