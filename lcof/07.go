package lcof

/**
 * 剑指 Offer 07. 重建二叉树
 * https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// WHY TO PURSUE THE INDEX, JUST AS FOLLOWS, LENGTH IS GREAT.
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	for i:=0; i<len(inorder); i++ {
//		if inorder[i] == preorder[0] {
//			return &TreeNode{
//				Val: preorder[0],
//				Left: buildTree(preorder[1:i+1], inorder[:i]),
//				Right: buildTree(preorder[i+1:], inorder[i+1:]),
//			}
//		}
//	}
//	return nil
//}

var rawPreorder, rawInorder []int
var preorderNodeMap, inorderNodeMap map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	nodeCnt := len(preorder)
	if nodeCnt == 0 {
		return nil
	}

	rawPreorder = make([]int, 0, nodeCnt)
	rawInorder = make([]int, 0, nodeCnt)
	preorderNodeMap = make(map[int]int, nodeCnt)
	inorderNodeMap = make(map[int]int, nodeCnt)
	for i := 0; i < nodeCnt; i++ {
		rawPreorder = append(rawPreorder, preorder[i])
		rawInorder = append(rawInorder, inorder[i])
		preorderNodeMap[preorder[i]] = i
		inorderNodeMap[inorder[i]] = i
	}

	return rebuildTree(preorder, inorder)
}

func rebuildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := preorder[0]
	t := &TreeNode{Val: root}

	// key point.
	inorderIdx := inorderNodeMap[root]
	var preorderIdx int
	for i := 0; i < inorderIdx; i++ {
		tPreorderIdx := preorderNodeMap[inorder[i]]
		if tPreorderIdx > preorderIdx {
			preorderIdx = tPreorderIdx
		}
	}

	// update map.
	rawInorderIdx := inorderNodeMap[root]
	delete(preorderNodeMap, root)
	delete(inorderNodeMap, root)
	rawPreorder = rawPreorder[1:]
	rawInorder = append(rawInorder[:rawInorderIdx], rawInorder[rawInorderIdx+1:]...)
	for i := 0; i < len(rawPreorder); i++ {
		preorderNodeMap[rawPreorder[i]]--
	}
	for i := rawInorderIdx; i < len(rawInorder); i++ {
		inorderNodeMap[rawInorder[i]]--
	}

	t.Left = rebuildTree(preorder[1:preorderIdx+1], inorder[:inorderIdx])
	t.Right = rebuildTree(preorder[preorderIdx+1:], inorder[inorderIdx+1:])

	return t
}
