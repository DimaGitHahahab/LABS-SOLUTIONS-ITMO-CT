package main

import "fmt"

func choose2num(choose []int, n, k int) {
	num := 0
	for i := 1; i <= k; i++ {
		for j := choose[i-1] + 1; j <= choose[i]-1; j++ {
			num += Cnk(n-j, k-i)
		}
	}
	fmt.Println(num)
	return
}

func Cnk(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	return n * Cnk(n-1, k-1) / k
}

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	choose := make([]int, k+1)
	for i := 1; i <= k; i++ {
		_, err = fmt.Scan(&choose[i])
		if err != nil {
			return
		}
	}
	choose[0] = 0
	choose2num(choose, n, k)
}
