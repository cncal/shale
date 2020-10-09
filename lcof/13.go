package lcof

/**
 * 剑指 Offer 13. 机器人的运动范围
 * https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
 */

func movingCount(m int, n int, k int) int {
	var cnt int
	temp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		temp = append(temp, make([]int, n))
	}

	// 当某个元素可以到达的时候，才可以将其上下左右的点入队
	queue := [][]int{{0, 0}}
	p := 0
	for p != len(queue) {
		x := queue[p][0]
		y := queue[p][1]
		if x >= 0 && x < m && y >= 0 && y < n && temp[x][y] == 0 {
			// 判断行坐标和列坐标的数位之和是否大于 K
			if checkSum(x, y, k) {
				temp[x][y] = -1 // 标识已经达到
				cnt++
				queue = append(
					queue,
					[]int{x - 1, y},
					[]int{x + 1, y},
					[]int{x, y - 1},
					[]int{x, y + 1},
				)
			}
		}
		p++
	}

	return cnt
}

func checkSum(x, y, k int) bool {
	var xs, ys int
	if x == 100 {
		xs = 1
	} else {
		xs = x%10 + x/10
	}
	if y == 100 {
		ys = 1
	} else {
		ys = y%10 + y/10
	}
	return xs+ys <= k
}
