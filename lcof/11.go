package lcof

/**
 * 剑指 Offer 11. 旋转数组的最小数字
 * https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
 */

func minArray(numbers []int) int {
	numCnt := len(numbers)
	if numCnt == 0 {
		return 0
	}
	for i := 1; i < numCnt; i++ {
		if numbers[i-1] > numbers[i] {
			return numbers[i]
		}
	}
	return numbers[0]
}
