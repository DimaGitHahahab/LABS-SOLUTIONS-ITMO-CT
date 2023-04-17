package main

import "fmt"

func splitting(n int, result []int) {
	if n == 0 {
		for _, v := range result[:len(result)-1] {
			fmt.Print(v, "+")
		}
		fmt.Println(result[len(result)-1])
		return
	}
	for i := 1; i <= n; i++ {
		if len(result) == 0 || i >= result[len(result)-1] {
			splitting(n-i, append(result, i))
		}
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	splitting(n, []int{})
}
