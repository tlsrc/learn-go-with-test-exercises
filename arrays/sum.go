package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) (sums []int) {
	for _, nums := range slices {
		sums = append(sums, Sum(nums))
	}
	return
}

func SumTails(slices ...[]int) []int {
	var sums = []int{}
	for _, nums := range slices {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}

	}
	return sums
}
