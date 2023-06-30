package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
)

type Node struct {
	left, right     *Node
	value, priority int
}

func (node *Node) exists(value int) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	if value < node.value {
		return node.left.exists(value)
	} else {
		return node.right.exists(value)
	}
}

func (node *Node) findMin() *Node {
	if node == nil || node.left == nil {
		return node
	}
	return node.left.findMin()
}

func (node *Node) findMax() *Node {
	if node == nil || node.right == nil {
		return node
	}
	return node.right.findMax()
}

func (node *Node) next(value int) *Node {
	if node == nil {
		return nil
	}
	if node.value <= value {
		return node.right.next(value)
	}
	res := node.left.next(value)
	if res == nil {
		return node
	}
	return res
}

func (node *Node) prev(value int) *Node {
	if node == nil {
		return nil
	}
	if node.value >= value {
		return node.left.prev(value)
	}
	res := node.right.prev(value)
	if res == nil {
		return node
	}
	return res
}

func (t *Node) split(x int) (*Node, *Node) {
	if t == nil {
		return nil, nil
	}
	if x < t.value {
		l, r := t.left.split(x)
		t.left = r
		return l, t
	} else {
		l, r := t.right.split(x)
		t.right = l
		return t, r
	}
}

func (l *Node) merge(r *Node) *Node {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.priority > r.priority {
		l.right = l.right.merge(r)
		return l
	} else {
		r.left = l.merge(r.left)
		return r
	}
}

func (t *Node) insert(x *Node) *Node {
	if t == nil {
		return x
	}
	if x.priority < t.priority {
		if x.value < t.value {
			t.left = t.left.insert(x)
		} else {
			t.right = t.right.insert(x)
		}
		return t
	}
	l, r := t.split(x.value)
	x.left = l
	x.right = r
	return x
}

func (t *Node) remove(x int) *Node {
	if t == nil {
		return nil
	}
	if x < t.value {
		t.left = t.left.remove(x)
		return t
	}
	if x > t.value {
		t.right = t.right.remove(x)
		return t
	}
	return t.left.merge(t.right)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	var root *Node
	var value int
	var command string

	for sc.Scan() {
		command = sc.Text()
		switch command {
		case "insert":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			if !root.exists(value) {
				node := &Node{nil, nil, value, rand.Int()}
				root = root.insert(node)
			}

		case "exists":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			if root.exists(value) {
				wr.WriteString("true\n")
			} else {
				wr.WriteString("false\n")
			}
		case "delete":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			if root.exists(value) {
				root = root.remove(value)
			}
		case "next":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			res := root.next(value)
			if res != nil {
				wr.WriteString(strconv.Itoa(res.value) + "\n")
			} else {
				wr.WriteString("none\n")
			}
		case "prev":
			sc.Scan()
			value, _ = strconv.Atoi(sc.Text())
			res := root.prev(value)
			if res != nil {
				wr.WriteString(strconv.Itoa(res.value) + "\n")
			} else {
				wr.WriteString("none\n")
			}
		}
	}
}
