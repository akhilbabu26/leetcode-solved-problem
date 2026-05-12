func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	remove := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {

		start := intervals[i][0]
		if start < end {
			remove++
		} else {
			end = intervals[i][1]
		}
	}
	return remove
}