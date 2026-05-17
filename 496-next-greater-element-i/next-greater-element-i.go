func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := []int{}
	nextGreater := make(map[int]int)

	for _, num := range nums2 {

		for len(stack) > 0 && num > stack[len(stack)-1] {

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			nextGreater[top] = num
		}

		stack = append(stack, num)
	}

	for len(stack) > 0 {

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nextGreater[top] = -1
	}

	result := []int{}

	for _, num := range nums1 {
		result = append(result, nextGreater[num])
	}

	return result
}