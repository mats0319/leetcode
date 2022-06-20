package mario

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeWrapper struct {
	node  *TreeNode
	times int
}

var nodeMap map[string]*nodeWrapper // tree desp - wrapper(node, appear times)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	nodeMap = make(map[string]*nodeWrapper) // for lc

	makeNodeMap(root)

	res := make([]*TreeNode, 0)
	for i := range nodeMap {
		item := nodeMap[i]

		if item.times > 1 {
			res = append(res, item.node)
		}
	}

	return res
}

func makeNodeMap(node *TreeNode) string {
	if node == nil {
		return ""
	}

	str := ""

	if node.Left == nil && node.Right == nil {
		str += strconv.Itoa(node.Val)
		wrapper, _ := nodeMap[str]
		if wrapper == nil {
			wrapper = &nodeWrapper{
				node:  node,
				times: 0,
			}
		}
		wrapper.times++
		nodeMap[str] = wrapper

		return str
	}

	leftChild := makeNodeMap(node.Left)
	rightChild := makeNodeMap(node.Right)

	str = strconv.Itoa(node.Val) + spliceChildren(leftChild, rightChild)
	wrapper, _ := nodeMap[str]
	if wrapper == nil {
		wrapper = &nodeWrapper{
			node:  node,
			times: 0,
		}
	}
	wrapper.times++
	nodeMap[str] = wrapper

	return str
}

func spliceChildren(left string, right string) string {
	if len(left)+len(right) < 1 {
		return ""
	}

	if len(right) > 0 {
		right = "," + right
	}

	return "(" + left + right + ")"
}
