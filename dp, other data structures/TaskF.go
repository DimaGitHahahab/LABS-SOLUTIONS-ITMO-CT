package main

import (
	"fmt"
	"math"
)

type stack struct {
	data []int
	top  int
}

func (s *stack) push(x int) {
	s.data = append(s.data, x)
	s.top++
}

func (s *stack) pop() int {
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.top--
	return x
}

// func to fill StackA at the beginning of the task
func (s *stack) leftPush(x int) {
	s.data = append([]int{x}, s.data...)
	s.top++
}

func main() {
	stackA := stack{[]int{}, -1}
	stackB := stack{[]int{}, -1}
	var ans []string
	var current int
	last := math.MinInt32
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	var y int
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&y)
		if err != nil {
			return
		}
		stackA.leftPush(y)
	}
	for stackA.top > -1 {
		if stackB.top == -1 {
			stackB.push(stackA.pop())
			ans = append(ans, "push")
		} else if stackA.data[stackA.top] < stackB.data[stackB.top] {
			stackB.push(stackA.pop())
			ans = append(ans, "push")
		} else {
			current = stackB.pop()
			if current < last {
				fmt.Println("impossible")
				return
			} else {
				ans = append(ans, "pop")
				last = current
			}
		}
	}
	for stackB.top > -1 {
		current = stackB.pop()
		if current < last {
			fmt.Println("impossible")
			return
		} else {
			ans = append(ans, "pop")
			last = current
		}
	}
	for _, v := range ans {
		fmt.Println(v)
	}
}
