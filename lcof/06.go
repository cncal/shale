package lcof

/**
 * 剑指 Offer 06. 从尾到头打印链表
 * https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
 */

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	t := make([]int, cnt)
	i := 0
	for head != nil {
		t[cnt-i-1] = head.Val
		head = head.Next
		i++
	}
	return t
}
