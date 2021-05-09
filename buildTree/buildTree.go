package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inOrderMap = make(map[int]int, 0)

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	r := buildTree(preorder, inorder)
	fmt.Println(r.Left.Val)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	for k, v := range inorder {
		inOrderMap[v] = k
	}
	return createTree(preorder, 0, 0, len(preorder)-1)
}

func createTree(preorder []int, start, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	rootInorderIndex := inOrderMap[preorder[start]]
	root := new(TreeNode)
	root.Val = preorder[start]
	root.Left = createTree(preorder, start+1, left, rootInorderIndex-1)
	root.Right = createTree(preorder, rootInorderIndex-left+start+1, rootInorderIndex+1, right)
	return root
}

//note
//通过跟结点在中序遍历中的位置，可以算出来，整个左子树的长度，和右子树的长度。
//在先序遍历中，跟结点的下一个是左子树的跟结点，加上左子树的长度之后，下一个是右子树的跟结点
//对于先序遍历和中序遍历来说，其右子树在数组中的位置是一致的
