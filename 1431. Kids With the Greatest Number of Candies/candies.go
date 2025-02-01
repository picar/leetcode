package candies

func maxValue(nums []int) int {
    if len(nums) == 0 {
        panic("slice must contain at least one element")
    }
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := maxValue(candies)
    result := make([]bool, len(candies))
    for i, c := range candies {
        result[i] = c+extraCandies >= max
    }
    return result
}