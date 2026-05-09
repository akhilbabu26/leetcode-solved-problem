import "sort"

func merge(intervals [][]int) [][]int {
   sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}

	for _, interval := range intervals {

		if len(result) == 0 ||
			result[len(result)-1][1] < interval[0] {

			result = append(result, interval)

		} else {
			last := result[len(result)-1]

			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		}
	}
	return result 
}