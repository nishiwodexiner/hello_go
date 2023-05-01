package main

import (
	"fmt"
)
//二叉树结构体定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//创建二叉树
func buildTree(values []int) *TreeNode {
	n := len(values)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue, idx := []*TreeNode{root}, 1

	for len(queue) > 0 && idx < n {
		node := queue[0]
		queue = queue[1:]

// 构建左子节点
		if idx < n && values[idx] != -1 { 
			node.Left = &TreeNode{Val: values[idx]}
			queue = append(queue, node.Left)
		}
		idx++

// 构建右子节点
		if idx < n && values[idx] != -1 { 
			node.Right = &TreeNode{Val: values[idx]}
			queue = append(queue, node.Right)
		}
		idx++
	}

	return root
}

//前序遍历
func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

//中序遍历
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%d ", root.Val)
	inorder(root.Right)
}

//后序遍历
func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Printf("%d ", root.Val)
}

func main() {
	fmt.Println("请输入层序遍历二叉树的节点值（使用 -1 表示空节点）：")
	var values []int
	var x int
	for {
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			break
		}
		values = append(values, x)
	}

	root := buildTree(values)


	fmt.Println("此二叉树前序遍历：")
	preorder(root)
	fmt.Println()

	fmt.Println("此二叉树中序遍历：")
	inorder(root)
	fmt.Println()

	fmt.Println("此二叉树后序遍历：")
	postorder(root)
}