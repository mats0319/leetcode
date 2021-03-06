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
	data    string
	nilFlag string
}

const nilFlag = "n"

func Constructor() Codec {
	return Codec{
		data:    "",
		nilFlag: nilFlag,
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	stack := make([]*TreeNode, 0, 1)
	stack = append(stack, root)

	data := ""
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		if node == nil {
			data += "," + nilFlag
			continue
		}

		data += "," + strconv.Itoa(node.Val)

		stack = append(stack, node.Left, node.Right)
	}

	return data[1:]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	payload := strings.Split(data, ",")
	if len(payload) == 1 && payload[0] == nilFlag {
		return nil
	}

	value, _ := strconv.Atoi(payload[0])
	root := &TreeNode{Val: value}
	stack := make([]*TreeNode, 0, 1)
	stack = append(stack, root)
	for i := 1; i < len(payload); i++ {
		node := stack[0]
		stack = stack[1:]

		if payload[i] != nilFlag {
			value, _ = strconv.Atoi(payload[i])
			node.Left = &TreeNode{Val: value}
			stack = append(stack, node.Left)
		}

		i++

		if payload[i] != nilFlag {
			value, _ = strconv.Atoi(payload[i])
			node.Right = &TreeNode{Val: value}
			stack = append(stack, node.Right)
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
