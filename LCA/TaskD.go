package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

type nodePair struct {
	dfsTime, index int
}

var dfsTime = 0

func findLowestCommonAncestor(firstNode, secondNode int, parentPowerMatrix [][]int, nodeDepth, parentNode []int) int {
	if nodeDepth[secondNode] < nodeDepth[firstNode] {
		firstNode, secondNode = secondNode, firstNode
	}
	depthDiff := nodeDepth[secondNode] - nodeDepth[firstNode]
	for power := 0; depthDiff > 0; power++ {
		if depthDiff%2 == 1 {
			secondNode = parentPowerMatrix[secondNode][power]
		}
		depthDiff /= 2
	}
	for parentNode[firstNode] != parentNode[secondNode] {
		power := 1
		for parentPowerMatrix[firstNode][power] != parentPowerMatrix[secondNode][power] {
			power++
		}
		firstNode = parentPowerMatrix[firstNode][power-1]
		secondNode = parentPowerMatrix[secondNode][power-1]
	}
	if secondNode == firstNode {
		return firstNode
	} else {
		return parentNode[firstNode]
	}
}

func depthFirstSearch(currentNode int, nodeDepth, dfsOrder, parentNode []int, adjacencyList [][]int, visited []bool) {
	dfsOrder[currentNode] = dfsTime
	dfsTime++
	visited[currentNode] = true
	for _, nextNode := range adjacencyList[currentNode] {
		nodeDepth[nextNode] = nodeDepth[currentNode] + 1
		parentNode[nextNode] = currentNode
		if !visited[nextNode] {
			depthFirstSearch(nextNode, nodeDepth, dfsOrder, parentNode, adjacencyList, visited)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var nodeCount, rootIndex int
	scanner.Scan()
	nodeCount, _ = strconv.Atoi(scanner.Text())

	adjacencyList := make([][]int, nodeCount)
	dfsOrder := make([]int, nodeCount)
	nodeDepth := make([]int, nodeCount)
	parentNode := make([]int, nodeCount)
	visited := make([]bool, nodeCount)
	parentPowerMatrix := make([][]int, nodeCount)
	for i := range parentPowerMatrix {
		parentPowerMatrix[i] = make([]int, int(math.Ceil(math.Log2(float64(nodeCount+1)))))
	}

	for i := 0; i < nodeCount; i++ {
		var parentNodeIndex int
		scanner.Scan()
		parentNodeIndex, _ = strconv.Atoi(scanner.Text())
		if parentNodeIndex == -1 {
			rootIndex = i
		} else {
			adjacencyList[parentNodeIndex-1] = append(adjacencyList[parentNodeIndex-1], i)
		}
	}

	depthFirstSearch(rootIndex, nodeDepth, dfsOrder, parentNode, adjacencyList, visited)

	for power := 0; power < len(parentPowerMatrix[0]); power++ {
		for node := 0; node < len(parentPowerMatrix); node++ {
			if power != 0 {
				parentPowerMatrix[node][power] = parentPowerMatrix[parentPowerMatrix[node][power-1]][power-1]
			} else {
				parentPowerMatrix[node][power] = parentNode[node]
			}
		}
	}

	var queryCount int
	scanner.Scan()
	queryCount, _ = strconv.Atoi(scanner.Text())

	for query := 0; query < queryCount; query++ {
		var nodeSubsetSize int

		scanner.Scan()
		nodeSubsetSize, _ = strconv.Atoi(scanner.Text())
		nodeSubset := make([]nodePair, nodeSubsetSize)
		for i := 0; i < nodeSubsetSize; i++ {
			var node int
			scanner.Scan()
			node, _ = strconv.Atoi(scanner.Text())
			nodeSubset[i] = nodePair{dfsOrder[node-1], node - 1}
		}

		sort.Slice(nodeSubset, func(i, j int) bool {
			return nodeSubset[i].dfsTime < nodeSubset[j].dfsTime
		})

		pathLength, previousNode := 0, 0
		for i := 0; i < nodeSubsetSize; i++ {
			if i != 0 {
				pathLength += nodeDepth[nodeSubset[i].index] - nodeDepth[findLowestCommonAncestor(previousNode, nodeSubset[i].index, parentPowerMatrix, nodeDepth, parentNode)]
			} else {
				pathLength += nodeDepth[nodeSubset[i].index]
			}
			previousNode = nodeSubset[i].index
		}

		writer.WriteString(strconv.Itoa(pathLength + 1))
		writer.WriteString("\n")
	}
}
