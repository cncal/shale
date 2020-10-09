package lcof

/**
 * 剑指 Offer 12. 矩阵中的路径
 * https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
 */

func exist(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		tmp := board[i][j]
		board[i][j] = '/' // HACK
		res := dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
		board[i][j] = tmp
		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
