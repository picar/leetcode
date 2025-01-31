package removeelement

func RemoveElement(nums []int, val int) int {
	k := len(nums) - 1
	for i := 0; i <= k; {
		if nums[i] != val {
			i++
		} else if nums[k] == val {
			k--
		} else if nums[i] == val && nums[k] != val {
			nums[i] = nums[k]
			i++
			k--
		}
	}
	return k + 1
}
