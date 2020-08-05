package backtracking

import (
	"fmt"
	"strconv"
)

/**
 * NO.401
 * https://leetcode-cn.com/problems/binary-watch/
 */
func readBinaryWatch(num int) []string {
	var results []string

	if num < 0 {
		return results
	}

	var backtrack func(int, []int, int, int, int)
	backtrack = func(num int, nums []int, start, hour, minute int) {
		if num == 0 {
			if hour > 11 || minute > 59 {
				return
			}

			var sMinute string
			if minute < 10 {
				sMinute = "0"
			}
			sMinute += strconv.Itoa(minute)

			results = append(
				results,
				fmt.Sprintf("%d:%s", hour, sMinute),
			)

			fmt.Println(fmt.Sprintf("%d:%s", hour, sMinute))
		}

		for i := start; i < len(nums); i++ {
			if i < 4 {
				hour += nums[i]
			} else {
				minute += nums[i]
			}

			backtrack(num-1, nums, i+1, hour, minute)

			if i < 4 {
				hour -= nums[i]
			} else {
				minute -= nums[i]
			}
		}
	}

	backtrack(num, []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}, 0, 0, 0)
	return results
}
