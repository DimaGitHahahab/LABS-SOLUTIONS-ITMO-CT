package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	in, _  = os.Open("problem2.in")
	sc     = bufio.NewScanner(in)
	out, _ = os.Create("problem2.out")
	wr     = bufio.NewWriter(out)
)

func ri() (v int) {
	sc.Scan()
	v, _ = strconv.Atoi(sc.Text())
	return
}

func rs() (v string) {
	sc.Scan()
	v = sc.Text()
	return
}

func wi(v int) {
	_, _ = wr.WriteString(strconv.Itoa(v))
}
func ws(v string) {
	_, _ = wr.WriteString(v)
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	word := rs()
	n, m, k := ri(), ri(), ri()
	terminals := make(map[int]bool)
	for i := 0; i < k; i++ {
		terminals[ri()-1] = true
	}
	automata := make([]map[string][]int, n)
	for i := 0; i < n; i++ {
		automata[i] = make(map[string][]int)
	}
	for i := 0; i < m; i++ {
		src, dst := ri()-1, ri()-1
		letter := rs()
		if _, ok := automata[src][letter]; ok {
			automata[src][letter] = append(automata[src][letter], dst)
		} else {
			automata[src][letter] = []int{dst}
		}
	}
	state := make(map[int]bool)
	state[0] = true
	for _, c := range word {
		newstate := make(map[int]bool)
		for s := range state {
			for _, dest := range automata[s][string(c)] {
				newstate[dest] = true
			}
		}
		state = newstate
	}
	var accepts bool
	for s := range state {
		if terminals[s] {
			accepts = true
		}
	}
	if accepts {
		ws("Accepts")
	} else {
		ws("Rejects")
	}
}
