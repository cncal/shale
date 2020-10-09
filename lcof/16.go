package lcof

/**
 * 剑指 Offer 16. 数值的整数次方
 * https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
 */
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickPow(x, n)
	} else {
		return 1.0 / quickPow(x, n*-1)
	}
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	y := quickPow(x, n/2)
	if n&1 == 1 {
		return y * y * x
	} else {
		return y * y
	}
}
