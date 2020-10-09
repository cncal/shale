package lcof

/**
 * 剑指 Offer 10- II. 青蛙跳台阶问题
 * https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
 */

func numWays(n int) int {
	if n == 0 {
		return 1
	} else if n < 3 {
		return n
	} else {
		x, y := 1, 2
		for i := 3; i <= n; i++ {
			x, y = y, x+y
			y %= 1e9 + 7
		}
		return y
	}
}
