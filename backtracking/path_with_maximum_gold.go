package backtracking

/**
 * NO.1219
 * https://leetcode-cn.com/problems/path-with-maximum-gold/comments/
 */

func getMaximumGold(grid [][]int) int {
	var max int

	xLen := len(grid)
	yLen := len(grid[0])

	var backtrack func([][]int, int, int, int, int)
	backtrack = func(grid [][]int, i, j, sum, cnt int) {
		if cnt == 25 {
			max = sum
			return
		}
		if sum > max {
			max = sum
		}

		t := grid[i][j]
		grid[i][j] = -1

		// upward
		if i > 0 {
			v := grid[i-1][j]
			if v > 0 {
				backtrack(grid, i-1, j, sum+v, cnt+1)
			}
		}

		// downward
		if i < xLen-1 {
			v := grid[i+1][j]
			if v > 0 {
				backtrack(grid, i+1, j, sum+v, cnt+1)
			}
		}

		// leftward
		if j > 0 {
			v := grid[i][j-1]
			if v > 0 {
				backtrack(grid, i, j-1, sum+v, cnt+1)
			}
		}

		// rightward
		if j < yLen-1 {
			v := grid[i][j+1]
			if v > 0 {
				backtrack(grid, i, j+1, sum+v, cnt+1)
			}
		}

		grid[i][j] = t
	}

	for i := 0; i < xLen; i++ {
		for j := 0; j < yLen; j++ {
			if grid[i][j] == 0 {
				continue
			}
			backtrack(grid, i, j, grid[i][j], 1)
		}
	}

	return max
}
