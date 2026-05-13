func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	n1 := len(nums1)
	n2 := len(nums2)

	low := 0
	high := n1

	for low <= high {
		partition1 := (low + high) / 2
		partition2 := (n1+n2+1)/2 - partition1
		left1 := math.MinInt
		if partition1 > 0 {
			left1 = nums1[partition1-1]
		}

		left2 := math.MinInt
		if partition2 > 0 {
			left2 = nums2[partition2-1]
		}
		right1 := math.MaxInt
		if partition1 < n1 {
			right1 = nums1[partition1]
		}

		right2 := math.MaxInt
		if partition2 < n2 {
			right2 = nums2[partition2]
		}
		if left1 <= right2 && left2 <= right1 {
			if (n1+n2)%2 == 0 {

				leftMax := max(left1, left2)
				rightMin := min(right1, right2)

				return float64(leftMax+rightMin) / 2.0
			}
			return float64(max(left1, left2))
		}
		if left1 > right2 {
			high = partition1 - 1
		} else {
			low = partition1 + 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}