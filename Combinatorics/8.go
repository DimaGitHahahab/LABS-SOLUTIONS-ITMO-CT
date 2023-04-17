package main

import "fmt"

func gen(p []int, n int, k int) {
	if len(p) == k {
		for i := 0; i < len(p); i++ {
			fmt.Print(p[i], " ")
		}
		fmt.Println()
		return
	}
	for c := 1; c <= n; c++ {
		if len(p) != 0 && c <= p[len(p)-1] {
			continue
		}
		if n-c < k-len(p)-1 {
			continue
		}
		gen(append(p, c), n, k)
	}
}

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	gen([]int{}, n, k)

}
