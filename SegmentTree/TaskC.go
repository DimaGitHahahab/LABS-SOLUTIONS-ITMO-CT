package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SegmentTree struct {
	tree []int
	arr  []int
}

func (s *SegmentTree) Set(i, v, x, xl, xr int) {
	if xl == xr-1 {
		s.tree[x] = v
		return
	}
	xm := (xl + xr) / 2
	if i < xm {
		s.Set(i, v, 2*x+1, xl, xm)
	} else {
		s.Set(i, v, 2*x+2, xm, xr)
	}
	s.tree[x] = s.tree[2*x+1] + s.tree[2*x+2]
}

func (s *SegmentTree) Sum(l, r, x, xl, xr int) int {
	if xl >= r || l >= xr {
		return 0
	}
	if xl >= l && xr <= r {
		return s.tree[x]
	}
	xm := (xl + xr) / 2
	sl := s.Sum(l, r, 2*x+1, xl, xm)
	sr := s.Sum(l, r, 2*x+2, xm, xr)
	return sl + sr
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		arr[i], _ = strconv.Atoi(sc.Text())
	}
	st := SegmentTree{
		tree: make([]int, 4*n),
		arr:  arr,
	}
	for i, v := range arr {
		st.Set(i, v, 0, 0, n)
	}
	for sc.Scan() {
		cmd := sc.Text()
		if cmd == "sum" {
			sc.Scan()
			l, _ := strconv.Atoi(sc.Text())
			sc.Scan()
			r, _ := strconv.Atoi(sc.Text())
			fmt.Println(strconv.Itoa(st.Sum(l-1, r, 0, 0, n)))
		} else if cmd == "set" {
			sc.Scan()
			i, _ := strconv.Atoi(sc.Text())
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			st.Set(i-1, v, 0, 0, n)
		}
	}
}
