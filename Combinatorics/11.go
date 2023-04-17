package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	var n int
	var s [][]int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		binary := strconv.FormatInt(int64(i), 2)
		for len(binary) < n {
			binary = "0" + binary
		}
		temp := make([]int, 0)
		for i := 0; i < n; i++ {
			if binary[i] == '1' {
				temp = append(temp, i+1)
			}
		}
		s = append(s, temp)
	}
	sort.Slice(s, func(i, j int) bool {
		for k := 0; k < len(s[i]) && k < len(s[j]); k++ {
			if s[i][k] != s[j][k] {
				return s[i][k] < s[j][k]
			}
		}
		return len(s[i]) < len(s[j])
	})
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			fmt.Print(s[i][j], " ")
		}
		fmt.Println()
	}

}
