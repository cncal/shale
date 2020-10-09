package lcof

/**
 * 剑指 Offer 10- I. 斐波那契数列
 * https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
 */

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		x, y := 0, 1
		for i := 2; i <= n; i++ {
			x, y = y, x+y
			y %= 1e9 + 7
		}
		return y
	}
}
