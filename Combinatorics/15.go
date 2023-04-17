package main

import (
	"fmt"
)

func getCombination(n, k, m int) []int {
	var choose []int
	next := 1
	for k > 0 {
		if m < Cnk(n-1, k-1) {
			choose = append(choose, next)
			k--
		} else {
			m -= Cnk(n-1, k-1)
		}
		n--
		next++
	}
	return choose
}

func Cnk(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	return n * Cnk(n-1, k-1) / k
}

func main() {
	var n, k, m int
	_, err := fmt.Scan(&n, &k, &m)
	if err != nil {
		return
	}
	ans := getCombination(n, k, m)
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
}
