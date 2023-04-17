package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func permutation2num(n int, a []int) {
	was := make([]bool, n+1)
	num := 0
	for i := 0; i < n; i++ {
		for j := 0; j < a[i]-1; j++ {
			if !was[j+1] {
				num += factorial(n - i - 1)
			}
		}
		was[a[i]] = true
	}
	fmt.Println(num)
	return
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&a[i])
		if err != nil {
			return
		}
	}
	permutation2num(n, a)
}
