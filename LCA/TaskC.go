package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxN = 50000
const maxLog = 20

var (
	g     [maxN + 1][]int
	cost  [maxN + 1][]int
	up    [maxN + 1][maxLog]int
	minUp [maxN + 1][maxLog]int
	h     [maxN + 1]int
)

func dfs(v, p int) {
	up[v][0] = p
	for i := 1; i < maxLog; i++ {
		up[v][i] = up[up[v][i-1]][i-1]
		minUp[v][i] = min(minUp[v][i-1], minUp[up[v][i-1]][i-1])
	}
	for _, to := range g[v] {
		if to != p {
			h[to] = h[v] + 1
			dfs(to, v)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lca(a, b int) int {
	if h[a] < h[b] {
		a, b = b, a
	}
	res := 1 << 30
	for i := maxLog - 1; i >= 0; i-- {
		if h[a]-(1<<i) >= h[b] {
			res = min(res, minUp[a][i])
			a = up[a][i]
		}
	}
	if a == b {
		return res
	}
	for i := maxLog - 1; i >= 0; i-- {
		if up[a][i] != up[b][i] {
			res = min(res, min(minUp[a][i], minUp[b][i]))
			a = up[a][i]
			b = up[b][i]
		}
	}
	return min(res, min(minUp[a][0], minUp[b][0]))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 1; i <= n; i++ {
		g[i] = make([]int, 0)
		cost[i] = make([]int, 0)
	}
	for i := 2; i <= n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		p, _ := strconv.Atoi(line[0])
		w, _ := strconv.Atoi(line[1])
		g[p] = append(g[p], i)
		g[i] = append(g[i], p)
		cost[i] = append(cost[i], w)
		cost[p] = append(cost[p], w)
		minUp[i][0] = w
	}
	dfs(1, 0)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		fmt.Println(lca(a, b))
	}
}
