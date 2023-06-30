package main

import (
	"bufio"
	"os"
	"strconv"
)

type Node struct {
	left, right *Node
	value       int
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{nil, nil, value}
	}
	if value < node.value {
		node.left = insert(node.left, value)
	} else {
		node.right = insert(node.right, value)
	}
	return node
}

func find(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	if value < node.value {
		return find(node.left, value)
	} else {
		return find(node.right, value)
	}
}

func erase(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = erase(node.left, value)
	} else if value > node.value {
		node.right = erase(node.right, value)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		}
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		minNode := findMin(node.right)
		node.value = minNode.value
		node.right = erase(node.right, minNode.value)
	}
	return node
}

func findMin(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return findMin(node.left)
}

func findMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return findMax(node.right)
}

func next(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if node.value <= value {
		return next(node.right, value)
	}
	res := next(node.left, value)
	if res == nil {
		return node
	}
	return res
}

func prev(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if node.value >= value {
		return prev(node.left, value)
	}
	res := prev(node.right, value)
	if res == nil {
		return node
	}
	return res
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	var (
		root    *Node
		value   int
		command string
	)
	for sc.Scan() {
		command = sc.Text()
		switch command {
		case "insert":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			if !find(root, value) {
				root = insert(root, value)
			}
		case "exists":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			if find(root, value) {
				wr.WriteString("true\n")
			} else {
				wr.WriteString("false\n")
			}
		case "delete":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			if find(root, value) {
				root = erase(root, value)
			}
		case "next":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			res := next(root, value)
			if res != nil {
				wr.WriteString(strconv.Itoa(res.value) + "\n")
			} else {
				wr.WriteString("none\n")
			}
		case "prev":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			res := prev(root, value)
			if res != nil {
				wr.WriteString(strconv.Itoa(res.value) + "\n")
			} else {
				wr.WriteString("none\n")
			}
		}
	}
}
