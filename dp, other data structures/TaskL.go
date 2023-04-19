package main

import "fmt"

func LIS(arr []int) (int, []int) {
	n := len(arr)
	prev := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
	}
	pos := 0
	length := dp[0]
	for i := 0; i < n; i++ {
		if dp[i] > length {
			pos = i
			length = dp[i]
		}
	}
	ans := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		ans[i] = arr[pos]
		pos = prev[pos]
	}
	return length, ans
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			return
		}
	}
	length, ans := LIS(arr)
	fmt.Println(length)
	for i := 0; i < length; i++ {
		fmt.Print(ans[i], " ")
	}
}
