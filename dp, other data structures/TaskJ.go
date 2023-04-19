package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	coins := make([]int, n, n)
	for i := 1; i < n-1; i++ {
		_, err := fmt.Scan(&coins[i])
		if err != nil {
			return
		}
	}
	maxCoins := make([]int, n, n)
	stepsForMaxCoins := make([]int, n, n)
	for i := 1; i < n; i++ {
		maxCoins[i] = math.MinInt32
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				maxCoins[i] = int(math.Max(float64(maxCoins[i]), float64(maxCoins[i-j]+coins[i])))
				if maxCoins[i] == maxCoins[i-j]+coins[i] {
					stepsForMaxCoins[i] = j
				}
			}
		}
	}
	fmt.Println(maxCoins[n-1])
	ans := make([]int, 0, n)
	for i := n - 1; i > 0; i -= stepsForMaxCoins[i] {
		ans = append(ans, i)
	}
	fmt.Println(len(ans))
	fmt.Print("1" + " ")
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Print(ans[i]+1, " ")
	}
}
