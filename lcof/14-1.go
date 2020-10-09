package lcof

/**
 * 剑指 Offer 14- I. 剪绳子
 * https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
 */

var nm map[int]int
var nmInit bool

func cuttingRope(n int) int {
	if !nmInit {
		nm = make(map[int]int, n)
		nm[1] = 1
		nm[2] = 1
		nmInit = true
	}
	var p int
	for i := n - 1; i > 0; i-- {
		np, ok := nm[n-i]
		if !ok {
			np = cuttingRope(n - i)
			nm[n-i] = np
		}
		if n-i > np {
			np = n - i
		}
		t := i * np
		if t > p {
			p = t
		}
	}
	return p
}
