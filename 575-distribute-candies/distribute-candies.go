func distributeCandies(candyType []int) int {
    toEat := len(candyType)/2
    candy := make(map[int]struct{})

    for _, v := range candyType{
        candy[v] = struct{}{}
        if len(candy) == toEat{
            return toEat
        }
    }
    return len(candy)
}