func intersect(nums1 []int, nums2 []int) []int {
   freq := make(map[int]int)

	for _, num := range nums1 {
		freq[num]++
	}

	var result []int

	for _, num := range nums2 {
		if freq[num] > 0 {
			result = append(result, num)
			freq[num]--
		}
	}
	return result 
}