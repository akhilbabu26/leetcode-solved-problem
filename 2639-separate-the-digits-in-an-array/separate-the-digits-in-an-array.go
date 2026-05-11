import "strconv"

func separateDigits(nums []int) []int {
    result := []int{}

    for i:= 0; i < len(nums); i++{
        s := strconv.Itoa(nums[i])
        num := 0
        for _, v := range s{
            if len(s) == 1{
                num = int(v)-'0'
                result = append(result, num)
                break
            }else{
                num = int(v)-'0'
                result = append(result, num)
            }
        }
    }

    
    return result
}