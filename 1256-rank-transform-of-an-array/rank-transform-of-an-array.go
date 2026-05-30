import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	sorted := make([]int, len(arr))
	copy(sorted, arr)

	sort.Ints(sorted)

	rankMap := make(map[int]int)
	rank := 1

	for _, num := range sorted {
		if _, exists := rankMap[num]; !exists {
			rankMap[num] = rank
			rank++
		}
	}

	result := make([]int, len(arr))

	for i, num := range arr {
		result[i] = rankMap[num]
	}

	return result
}