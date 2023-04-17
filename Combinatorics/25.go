package main

import "fmt"

func nextChoose(a []int, n, k int) ([]int, bool) {
	b := make([]int, k+1)
	for i := 0; i < k; i++ {
		b[i] = a[i]
	}
	b[k] = n + 1
	i := k - 1
	for i >= 0 && b[i+1]-b[i] < 2 {
		i--
	}
	if i >= 0 {
		b[i]++
		for j := i + 1; j < k; j++ {
			b[j] = b[j-1] + 1
		}
		for i := 0; i < k; i++ {
			a[i] = b[i]
		}
		return a, true
	} else {
		return nil, false
	}
}

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	a := make([]int, k)
	for i := 0; i < k; i++ {
		_, err := fmt.Scan(&a[i])
		if err != nil {
			return
		}
	}
	ans, ok := nextChoose(a, n, k)
	if !ok {
		fmt.Println(-1)
	} else {
		for i := 0; i < k; i++ {
			fmt.Print(ans[i], " ")
		}
	}

}