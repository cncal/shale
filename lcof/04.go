package lcof

/**
 * 剑指 Offer 04. 二维数组中的查找
 * https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
 */

func findNumberIn2DArray(matrix [][]int, target int) bool {
	hL := len(matrix)
	if hL == 0 {
		return false
	}
	vL := len(matrix[0])
	if vL == 0 {
		return false
	}
	if target < matrix[0][0] {
		return false
	}
	if target > matrix[hL-1][vL-1] {
		return false
	}
	for i := 0; i < hL; i++ {
		for j := 0; j < vL; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
