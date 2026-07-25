import "strconv"

func maxProduct(n int) int {
    s := strconv.Itoa(n)

    max1, max2 := 0, 0

    for _, ch := range s {
        digit := int(ch - '0')

        if digit >= max1 {
            max2 = max1
            max1 = digit
        } else if digit > max2 {
            max2 = digit
        }
    }

    return max1 * max2
}