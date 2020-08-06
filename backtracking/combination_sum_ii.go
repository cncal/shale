package backtracking

import (
	"sort"
)

/**
 * NO.40
 * https://leetcode-cn.com/problems/combination-sum-ii/
 */
func combinationSum2(candidates []int, target int) [][]int {
	var set [][]int

	if len(candidates) == 0 {
		return set
	}

	sort.Ints(candidates)

	var backtrack func(int, []int, int, []int, int)
	backtrack = func(start int, candidates []int, target int, solution []int, sum int) {
		if sum == target {
			temp := make([]int, len(solution))
			copy(temp, solution)
			set = append(set, temp)
			return
		}

		for i := start; i < len(candidates); i++ {
			// pruning.
			if sum+candidates[i] > target {
				break
			}
			// tips.
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// make choice.
			solution = append(solution, candidates[i])

			// backtrack.
			backtrack(i+1, candidates, target, solution, sum+candidates[i])

			// cancel.
			preSolution := solution[:len(solution)-1]
			solution = preSolution
		}
	}

	backtrack(0, candidates, target, []int{}, 0)

	return set
}
