func getSumAbsoluteDifferences(nums []int) []int {
    n := len(nums)
	result := make([]int, n)
	total := 0
	for _, v := range nums {
		total += v
	}
	leftSum := 0
	for i := 0; i < n; i++ {
		rightSum := total - leftSum - nums[i]

		left := nums[i]*i - leftSum

		right := rightSum - nums[i]*(n-i-1)

		result[i] = left + right

		leftSum += nums[i]
	}

	return result
}