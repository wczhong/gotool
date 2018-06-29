package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Data int
	N0   *Node
	N1   *Node
	N2   *Node
	N3   *Node
	N4   *Node
	N5   *Node
	N6   *Node
	N7   *Node
	N8   *Node
	N9   *Node
}

func StructDictTree(head *Node, data []int) {
	if head == nil || len(data) == 0 {
		return
	}

	for _, d := range data {
		root := head
		AddDataToTree(root, d)
	}
}

func AddDataToTree(head *Node, data int) {
	dataStr := strconv.Itoa(data)
	for index, vStr := range dataStr {
		vInt, _ := strconv.Atoi(string(vStr))
		flag := 0
		if index == len(dataStr)-1 {
			flag = 1
		}
		head = AddNodeToTree(head, vInt, flag, data)
	}
}

func AddNodeToTree(head *Node, key int, flag int, data int) *Node {
	// fmt.Println(key)
	node := &Node{}
	switch key {
	case 0:
		if head.N0 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N0 = n
			node = n
		} else {
			node = head.N0
		}
	case 1:
		if head.N1 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N1 = n
			node = n
		} else {
			node = head.N1
		}
	case 2:
		if head.N2 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N2 = n
			node = n
		} else {
			node = head.N2
		}
	case 3:
		if head.N3 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N3 = n
			node = n
		} else {
			node = head.N3
		}
	case 4:
		if head.N4 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N4 = n
			node = n
		} else {
			node = head.N4
		}
	case 5:
		if head.N5 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N5 = n
			node = n
		} else {
			node = head.N5
		}
	case 6:
		if head.N6 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N6 = n
			node = n
		} else {
			node = head.N6
		}
	case 7:
		if head.N7 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N7 = n
			node = n
		} else {
			node = head.N7
		}
	case 8:
		if head.N8 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N8 = n
			node = n
		} else {
			node = head.N8
		}
	case 9:
		if head.N9 == nil {
			n := &Node{
				Data: -1,
			}
			if flag == 1 {
				n.Data = data
			}
			head.N9 = n
			node = n
		} else {
			node = head.N9
		}
	}
	return node
}

func IteralTree(root *Node) {
	// fmt.Println(root)
	if root.Data != -1 {
		fmt.Println(root.Data)
	}
	if root.N0 != nil {
		IteralTree(root.N0)
	}
	if root.N1 != nil {
		IteralTree(root.N1)
	}
	if root.N2 != nil {
		IteralTree(root.N2)
	}
	if root.N3 != nil {
		IteralTree(root.N3)
	}
	if root.N4 != nil {
		IteralTree(root.N4)
	}
	if root.N5 != nil {
		IteralTree(root.N5)
	}
	if root.N6 != nil {
		IteralTree(root.N6)
	}
	if root.N7 != nil {
		IteralTree(root.N7)
	}
	if root.N8 != nil {
		IteralTree(root.N8)
	}
	if root.N9 != nil {
		IteralTree(root.N9)
	}

}
func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 100, 101, 102, 201, 202, 303, 405, 504, 601, 706, 801, 909, 10101, 45, 78, 90, 1, 23, 56, 90, 0}
	root := &Node{
		Data: -1,
	}
	StructDictTree(root, data)
	IteralTree(root)
}
