func totalWaviness(num1 int, num2 int) int {
	ans := 0

	for num := num1; num <= num2; num++ {
		ans += waviness(num)
	}

	return ans
}

func waviness(num int) int {
	digits := []int{}

	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if len(digits) < 3 {
		return 0
	}

	cnt := 0

	for i := 1; i < len(digits)-1; i++ {
		if digits[i] > digits[i-1] && digits[i] > digits[i+1] {
			cnt++
		} else if digits[i] < digits[i-1] && digits[i] < digits[i+1] {
			cnt++
		}
	}

	return cnt
}