package main

import "fmt"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func gen(n int, p []int) {
	if len(p) == n {
		for _, v := range p {
			fmt.Print(v, " ")
		}
		fmt.Println()
		return
	}
	for i := 1; i <= n; i++ {
		if !contains(p, i) {
			gen(n, append(p, i))
		}
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Shit")
	}
	gen(n, []int{})
}
