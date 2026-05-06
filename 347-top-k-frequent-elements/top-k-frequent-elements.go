func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)

    for _, v := range nums{
        freq[v]++
    }

    bucket := make([][]int, len(nums)+1)

    for num, count := range freq{
        bucket[count] = append(bucket[count], num)
    }

    result := []int{}

    for i := len(bucket)-1; i>0; i--{
        for _, v := range bucket[i]{
            result = append(result, v)

            if len(result) == k{
                return result
            }
        }
    }

    return result
}