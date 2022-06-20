package mario

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	sep     string
	nilNode string
}

func Constructor() Codec {
	return Codec{
		sep:     ",",
		nilNode: "#",
	}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	res := ""

	layer := []*TreeNode{root}
	for len(layer) > 0 {
		newLayer := make([]*TreeNode, 0, len(layer)*2)

		for len(layer) > 0 {
			node := layer[0]
			layer = layer[1:]

			if node == nil {
				res += c.nilNode + c.sep
				continue
			}

			res += strconv.Itoa(node.Val) + c.sep
			newLayer = append(newLayer, node.Left, node.Right)
		}

		layer = newLayer
	}

	return res[:len(res)-1]
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) < 1 {
		return nil
	}

	valueStrSlice := strings.Split(data, c.sep)
	val, _ := strconv.Atoi(valueStrSlice[0])
	root := &TreeNode{Val: val}

	layer := []*TreeNode{root}
	index := 1
	for len(layer) > 0 {
		newLayer := make([]*TreeNode, 0, len(layer)*2)

		for len(layer) > 0 {
			node := layer[0]
			layer = layer[1:]

			leftValStr := valueStrSlice[index]
			rightValStr := valueStrSlice[index+1]

			index += 2

			if leftValStr == c.nilNode {
				node.Left = nil
			} else {
				val, _ = strconv.Atoi(leftValStr)
				node.Left = &TreeNode{Val: val}
				newLayer = append(newLayer, node.Left)
			}

			if rightValStr == c.nilNode {
				node.Right = nil
			} else {
				val, _ = strconv.Atoi(rightValStr)
				node.Right = &TreeNode{Val: val}
				newLayer = append(newLayer, node.Right)
			}
		}

		layer = newLayer
	}

	return root
}
