package main

import "fmt"

func balls(a []int) int {
	ans := 0
	var count int
	for i := 0; i < len(a); i++ {
		count = 0
		j := i
		for j < len(a) && (a[j] == a[i] || (a[j] == -1 && a[i] != -1)) {
			if a[j] == -1 {
				count++
			}
			j++
		}
		if j-i-count >= 3 {
			ans += j - i - count
			for k := i; k < j; k++ {
				a[k] = -1
			}
			i = -1
		} else {
			i = j - 1
		}

	}
	return ans
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	chain := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&chain[i])
		if err != nil {
			return
		}
	}
	fmt.Println(balls(chain))
}
