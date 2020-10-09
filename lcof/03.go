package lcof

/**
 * 剑指 Offer 03. 数组中重复的数字
 * https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
 */

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := m[num]; ok {
			return num
		} else {
			m[num] = struct{}{}
		}
	}
	return -1
}
