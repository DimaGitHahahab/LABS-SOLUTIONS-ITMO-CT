package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func getPermutations(n, k int) []int {
	permutation := make([]int, n)
	was := make([]bool, n)
	for i := 0; i < n; i++ {
		alreadyWas := k / factorial(n-i-1)
		k = k % factorial(n-i-1)
		curFree := 0
		for j := 0; j < n; j++ {
			if !was[j] {
				curFree++
				if curFree == alreadyWas+1 {
					permutation[i] = j
					was[j] = true
				}
			}
		}
	}
	return permutation
}

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	permutation := getPermutations(n, k)
	for i := 0; i < n; i++ {
		fmt.Print(permutation[i]+1, " ")
	}
}
