func minElement(nums []int) int {
   minSum := int(^uint(0) >> 1)

	for _, num := range nums {
		sum := 0

		for num > 0 {
			sum += num % 10
			num /= 10
		}

		if sum < minSum {
			minSum = sum
		}
	}

	return minSum 
}