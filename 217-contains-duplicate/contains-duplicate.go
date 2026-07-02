func containsDuplicate(nums []int) bool {
    m := make(map[int]int, len(nums))

    for _, v := range nums{
        m[v]++
        if m[v] > 1{
            return true
        }
    }

    return false
}