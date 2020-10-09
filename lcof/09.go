package lcof

/**
 * 剑指 Offer 09. 用两个栈实现队列
 * https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
 */
type CQueue struct {
	stack1 []int // 仅支持插入
	stack2 []int // 仅支持删除
}

func Constructor() CQueue {
	return CQueue{stack1: nil, stack2: nil}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	stack2cnt := len(this.stack2)
	if len(this.stack2) > 0 {
		head := this.stack2[stack2cnt-1]
		this.stack2 = this.stack2[:stack2cnt-1]
		return head
	}

	stack1cnt := len(this.stack1)
	if stack1cnt == 0 {
		return -1
	}
	head := this.stack1[0]
	this.stack2 = make([]int, 0, stack1cnt-1)
	for i := stack1cnt - 1; i > 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	this.stack1 = nil
	return head
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
