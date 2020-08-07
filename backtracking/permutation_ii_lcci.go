package backtracking

import (
	"strings"
)

/**
 * 面试题 08.08
 * https://leetcode-cn.com/problems/permutation-ii-lcci/
 *
 * TODO: learn to prune.
 */

func permutation(S string) []string {
	if len(S) == 1 {
		return []string{S}
	}

	results := make([]string, 0)
	resultMap := make(map[string]bool)
	placeholder := " "

	var backtrack func([]string, string)
	backtrack = func(rs []string, s string) {
		var placeholderCnt int
		for i := 0; i < len(rs); i++ {
			if rs[i] == placeholder {
				placeholderCnt++

				if placeholderCnt == len(rs) {
					if _, ok := resultMap[s]; !ok {
						results = append(results, s)
						resultMap[s] = true
					}
					return
				}
				continue
			}

			c := rs[i]
			s += c
			rs[i] = placeholder

			backtrack(rs, s)

			s = s[:len(s)-1]
			rs[i] = c
		}
	}

	rs := strings.Split(S, "")
	backtrack(rs, "")

	return results
}
