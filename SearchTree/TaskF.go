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
	size            int
}

func (node *Node) updateSize() {
	node.size = 1
	if node.left != nil {
		node.size += node.left.size
	}
	if node.right != nil {
		node.size += node.right.size
	}
}

func (t *Node) kthMaximum(k int) *Node {
	if t == nil {
		return nil
	}
	rightSize := 0
	if t.right != nil {
		rightSize = t.right.size
	}
	if k == rightSize+1 {
		return t
	}
	if k <= rightSize {
		return t.right.kthMaximum(k)
	}
	return t.left.kthMaximum(k - rightSize - 1)
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
	var l, r *Node
	if x < t.value {
		l, r = t.left.split(x)
		t.left = r
		t.updateSize()
		return l, t
	} else {
		l, r = t.right.split(x)
		t.right = l
		t.updateSize()
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
		l.updateSize()
		return l
	} else {
		r.left = l.merge(r.left)
		r.updateSize()
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
		t.updateSize()
		return t
	}
	l, r := t.split(x.value)
	x.left = l
	x.right = r
	x.updateSize()
	return x
}

func (t *Node) remove(x int) *Node {
	if t == nil {
		return nil
	}
	if x < t.value {
		t.left = t.left.remove(x)
	} else if x > t.value {
		t.right = t.right.remove(x)
	} else {
		t = t.left.merge(t.right)
	}
	if t != nil {
		t.updateSize()
	}
	return t
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	var root *Node
	var k int
	var command string
	sc.Scan()
	for sc.Scan() {
		command = sc.Text()
		sc.Scan()
		k, _ = strconv.Atoi(sc.Text())
		switch command {
		case "1":
			root = root.insert(&Node{nil, nil, k, rand.Int(), 1})
		case "0":
			kthMax := root.kthMaximum(k)
			if kthMax != nil {
				wr.WriteString(strconv.Itoa(kthMax.value) + "\n")
			}
		case "-1":
			root = root.remove(k)
		}
	}
}
