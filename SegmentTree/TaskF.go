package main

import (
	"fmt"
	"math"
)

type SparseTable struct {
	lookup [][]int
}

func NewSparseTable(nums []int) *SparseTable {
	n := len(nums)
	k := int(math.Log2(float64(n))) + 1

	// Create lookup table
	lookup := make([][]int, k)
	for i := 0; i < k; i++ {
		lookup[i] = make([]int, n-(1<<i)+1)
		for j := 0; j < n-(1<<i)+1; j++ {
			if i == 0 {
				lookup[i][j] = nums[j]
			} else {
				lookup[i][j] = min(lookup[i-1][j], lookup[i-1][j+(1<<(i-1))])
			}
		}
	}

	return &SparseTable{lookup: lookup}
}

func (st *SparseTable) Query(l, r int) int {
	k := int(math.Log2(float64(r - l + 1)))
	return min(st.lookup[k][l], st.lookup[k][r-(1<<k)+1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, m, a0 int
	fmt.Scan(&n, &m, &a0)
	var u, v int
	fmt.Scan(&u, &v)
	a := make([]int, n)
	a[0] = a0
	for i := 1; i < n; i++ {
		a[i] = (23*a[i-1] + 21563) % 16714589
	}
	st := NewSparseTable(a)
	var uNext, vNext, tempResult int
	tempResult = st.Query(u-1, v-1)
	for i := 1; i < m; i++ {
		uNext = ((17*u + 751 + tempResult + 2*i) % n) + 1
		vNext = ((13*v + 593 + tempResult + 5*i) % n) + 1
		var left, right int
		if uNext < vNext {
			left = uNext
			right = vNext
		} else {
			left = vNext
			right = uNext
		}
		tempResult = st.Query(left-1, right-1)
		u, v = uNext, vNext
	}
	if uNext == 0 && vNext == 0 {
		uNext = u
		vNext = v
	}
	fmt.Println(uNext, vNext, tempResult)
}
