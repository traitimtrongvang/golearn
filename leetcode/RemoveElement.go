package leetcode

// 27. Remove Element

func RemoveElement(nums []int, val int) int {

	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}

	return flag
}
