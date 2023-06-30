package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	currentNode      = 1
	logNumberOfNodes = 0
	numberOfNodes    int
	depth            []int
	dp               [][]int
	parent           []int
)

func addNode(parentNodeIndex int) {
	parent[currentNode] = currentNode
	dp[currentNode][0] = parentNodeIndex
	depth[currentNode] = 1 + depth[parentNodeIndex]

	for j := 1; 1<<(j-1) < numberOfNodes; j++ {
		dp[currentNode][j] = dp[dp[currentNode][j-1]][j-1]
	}

	currentNode++
}

func findRoot(x int) int {
	if parent[x] != x {
		parent[x] = findRoot(parent[x])
	}
	return parent[x]
}

func merge(x, y int) {
	rootOfX := findRoot(x)
	rootOfY := findRoot(y)

	if rootOfX != rootOfY {
		parent[rootOfX] = rootOfY
	}
}

func lowestCommonAncestor(u, v int) int {
	if depth[v] > depth[u] {
		u, v = v, u
	}

	for i := logNumberOfNodes; i >= 0; i-- {
		if depth[u]-(1<<i) >= depth[v] {
			u = dp[u][i]
		}
	}

	if u == v {
		return u
	}

	for i := logNumberOfNodes; i >= 0; i-- {
		if dp[v][i] != dp[u][i] {
			v = dp[v][i]
			u = dp[u][i]
		}
	}

	return dp[u][0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	numberOfNodes, _ = strconv.Atoi(scanner.Text())

	depth = make([]int, numberOfNodes)
	dp = make([][]int, numberOfNodes)
	parent = make([]int, numberOfNodes)

	for j := 1; 1<<(j-1) < numberOfNodes; j++ {
		logNumberOfNodes++
	}

	for i := 0; i < numberOfNodes; i++ {
		dp[i] = make([]int, logNumberOfNodes+1)
	}

	for i := 0; i < numberOfNodes; i++ {
		scanner.Scan()
		line := scanner.Text()
		command := line[0]
		switch command {
		case '+':
			parentNodeIndex, _ := strconv.Atoi(strings.Trim(line[1:], " "))
			addNode(parentNodeIndex - 1)
		case '-':
			nodeIndex, _ := strconv.Atoi(strings.Trim(line[1:], " "))
			merge(nodeIndex-1, dp[nodeIndex-1][0])
		case '?':
			parts := strings.Split(strings.Trim(line[1:], " "), " ")
			u, _ := strconv.Atoi(parts[0])
			v, _ := strconv.Atoi(parts[1])
			ans := findRoot(lowestCommonAncestor(u-1, v-1)) + 1
			stringAns := strconv.Itoa(ans)
			writer.WriteString(stringAns)
			writer.WriteString("\n")
		}
	}
}
