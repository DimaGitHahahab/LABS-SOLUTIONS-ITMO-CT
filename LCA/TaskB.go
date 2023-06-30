package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	n, m, LOG int
	g         [][]int
	up        [][]int
	h         []int
)

func dfs(v, p int) {
	up[v][0] = p
	for i := 1; i <= LOG; i++ {
		up[v][i] = up[up[v][i-1]][i-1]
	}
	for _, u := range g[v] {
		if u != p {
			h[u] = h[v] + 1
			dfs(u, v)
		}
	}
}

func lca(u, v int) int {
	if h[u] < h[v] {
		u, v = v, u
	}
	for i := LOG; i >= 0; i-- {
		if h[u]-(1<<i) >= h[v] {
			u = up[u][i]
		}
	}
	if u == v {
		return u
	}
	for i := LOG; i >= 0; i-- {
		if up[u][i] != up[v][i] {
			u = up[u][i]
			v = up[v][i]
		}
	}
	return up[u][0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	LOG = 1
	for (1 << LOG) <= n {
		LOG++
	}
	g = make([][]int, n+1)
	up = make([][]int, n+1)
	for i := range up {
		up[i] = make([]int, LOG+1)
	}
	h = make([]int, n+1)
	for i := 2; i <= n; i++ {
		scanner.Scan()
		p, _ := strconv.Atoi(scanner.Text())
		g[p] = append(g[p], i)
	}
	dfs(1, 1)
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < m; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		fmt.Println(lca(u, v))
	}
}
