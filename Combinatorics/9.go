package main

import "fmt"

func gen(n, cOpen, cClose int, ans string) {
	if cOpen+cClose == 2*n {
		fmt.Println(ans)
		return
	}
	if cOpen < n {
		gen(n, cOpen+1, cClose, ans+"(")
	}
	if cOpen > cClose {
		gen(n, cOpen, cClose+1, ans+")")
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	gen(n, 0, 0, "")
}
