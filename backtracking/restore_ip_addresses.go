package backtracking

import (
	"strconv"
)

/**
 * NO.93
 * https://leetcode-cn.com/problems/restore-ip-addresses/
 */

const SegCnt = 4

var (
	ans      []string
	segments []int
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, SegCnt)
	ans = []string{}
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, segId, segStart int) {
	// 如果找到了 4 段 IP 地址并且遍历完了字符串，那么就是一种答案
	if segId == SegCnt {
		if segStart == len(s) {
			ipAddr := ""
			for i := 0; i < SegCnt; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != SegCnt-1 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		return
	}

	// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
	if segStart == len(s) {
		return
	}
	// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[segStart] == '0' {
		segments[segId] = 0
		dfs(s, segId+1, segStart+1)
	}
	// 一般情况，枚举每一种可能性并递归
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[segId] = addr
			dfs(s, segId+1, segEnd+1)
		} else {
			break
		}
	}
}

//func restoreIpAddresses(s string) []string {
//	var results []string
//
//	sl := len(s)
//	if sl < 4 {
//		return results
//	} else if sl == 4 {
//		results = []string{strings.Join(strings.Split(s, ""), ".")}
//		return results
//	}
//
//	m := make(map[string]bool)
//
//	var backtrack func(string, []int, int)
//	backtrack = func(s string, dotIdx []int, start int) {
//		if len(dotIdx) == 3 {
//			if start <= len(s)-1 {
//				if s[start] == '0' && start < len(s)-1 {
//					return
//				}
//				t := s[start:]
//				if it, err := strconv.Atoi(t); err == nil && it <= 255 {
//					address := fmt.Sprintf(
//						"%s.%s.%s.%s",
//						s[0:dotIdx[0]],
//						s[dotIdx[0]:dotIdx[1]],
//						s[dotIdx[1]:dotIdx[2]],
//						s[dotIdx[2]:],
//					)
//					if _, ok := m[address]; !ok {
//						results = append(
//							results,
//							address,
//						)
//						m[address] = true
//					}
//				}
//			}
//			return
//		}
//
//		for i := start; i < len(s); i++ {
//			for j := 1; j < 4; j++ {
//				if i+j < len(s) {
//					if t, err := strconv.Atoi(s[start : i+j]); err == nil && t <= 255 {
//						dotIdx = append(dotIdx, i+j)
//						backtrack(s, dotIdx, i+j)
//						if s[i] != '0' {
//							dotIdx = dotIdx[:len(dotIdx)-1]
//						}
//					}
//				}
//			}
//		}
//	}
//
//	backtrack(s, []int{}, 0)
//
//	return results
//}
