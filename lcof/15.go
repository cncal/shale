package lcof

/**
 * 剑指 Offer 15. 二进制中1的个数
 * https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
 */

func hammingWeight(num uint32) int {
	var cnt int
	for num != 0 {
		cnt++
		num &= num - 1
	}
	return cnt
}
