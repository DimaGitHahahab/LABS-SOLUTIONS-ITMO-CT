package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value    int64
	priority int
	count    int
	sum      int64
	left     *Node
	right    *Node
}

type TreapWithSums struct {
	root *Node
}

type Pair struct {
	first  *Node
	second *Node
}

func (t *TreapWithSums) generatePriority() int {
	return 1000 + rand.Intn(10000)
}

func (t *TreapWithSums) calculateSum(cur *Node) int64 {
	if cur == nil {
		return 0
	}
	return cur.sum
}

func (t *TreapWithSums) updateCount(r *Node) {
	if r == nil {
		return
	}
	leftCount := 0
	if r.left != nil {
		leftCount = r.left.count
	}
	rightCount := 0
	if r.right != nil {
		rightCount = r.right.count
	}
	r.count = leftCount + rightCount + 1
	r.sum = t.calculateSum(r.left) + t.calculateSum(r.right) + r.value
}

func (t *TreapWithSums) split(r *Node, x int64) Pair {
	if r == nil {
		return Pair{nil, nil}
	}
	if r.value <= x {
		p := t.split(r.right, x)
		r.right = p.first
		t.updateCount(r)
		return Pair{r, p.second}
	}
	p := t.split(r.left, x)
	r.left = p.second
	t.updateCount(r)
	return Pair{p.first, r}
}

func (t *TreapWithSums) merge(r1 *Node, r2 *Node) *Node {
	if r1 == nil {
		return r2
	}
	if r2 == nil {
		return r1
	}
	if r1.priority > r2.priority {
		r1.right = t.merge(r1.right, r2)
		t.updateCount(r1)
		return r1
	}
	r2.left = t.merge(r1, r2.left)
	t.updateCount(r2)
	return r2
}

func (t *TreapWithSums) add(x int64) {
	newNode := &Node{value: x, priority: t.generatePriority(), count: 1, sum: x}
	p1 := t.split(t.root, x-1)
	p2 := t.split(p1.second, x)
	t.root = t.merge(p1.first, t.merge(newNode, p2.second))
}

func (t *TreapWithSums) sum(l int64, r int64) int64 {
	p1 := t.split(t.root, l-1)
	p2 := t.split(p1.second, r)
	ans := int64(0)
	if p2.first != nil {
		ans = p2.first.sum
	}
	t.root = t.merge(p1.first, t.merge(p2.first, p2.second))
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	var lastQueryResult int64
	const mod int64 = 1_000_000_000
	var ans strings.Builder
	treap := TreapWithSums{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		if line[0] == "+" {
			num, _ := strconv.ParseInt(line[1], 10, 64)
			treap.add((lastQueryResult + num) % mod)
			lastQueryResult = 0
		} else {
			l, _ := strconv.ParseInt(line[1], 10, 64)
			r, _ := strconv.ParseInt(line[2], 10, 64)
			answer := treap.sum(l, r)
			ans.WriteString(fmt.Sprintf("%d\n", answer))
			lastQueryResult = answer
		}
	}
	fmt.Print(ans.String())
}
