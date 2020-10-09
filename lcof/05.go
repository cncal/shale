package lcof

import (
	"bytes"
)

/**
 * 剑指 Offer 05. 替换空格
 * https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
 */

func replaceSpace(s string) string {
	sl := len(s)
	t := make([]byte, 0, sl*3)
	size := 0
	for i := 0; i < sl; i++ {
		if s[i] == ' ' {
			t = append(t, '%', '2', '0')
			size += 3
		} else {
			t = append(t, s[i])
			size++
		}
	}
	buf := bytes.NewBuffer(t[:size])
	return buf.String()
}
