package arrays

func Sum(numbers [5]int) (sum int) {
	for _, n := range numbers {
		sum += n
	}

	return
}

func SumSlice(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}

	return
}

func SumAll(numbers ...[]int) []int {
	// sum := make([]int, len(numbers))
	// for i, nums := range numbers {
	// 	sum[i] = SumSlice(nums)
	// }
	
	var sum []int
	for _, nums := range numbers {
		sum = append(sum, SumSlice(nums))
	}

	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SumSlice(nums[1:]))
		}
	}
	return sums
}