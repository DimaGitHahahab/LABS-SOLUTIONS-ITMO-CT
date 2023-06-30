package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	in, _  = os.Open("problem1.in")
	sc     = bufio.NewScanner(in)
	out, _ = os.Create("problem1.out")
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
	automata := make([]map[string]int, n)
	for i := 0; i < n; i++ {
		automata[i] = make(map[string]int)
	}
	for i := 0; i < m; i++ {
		src, dst := ri()-1, ri()-1
		automata[src][rs()] = dst
	}
	var (
		state int
		ok    bool
	)
	for _, c := range word {
		state, ok = automata[state][string(c)]
		if !ok {
			ws("Rejects")
			return
		}
	}
	if terminals[state] {
		ws("Accepts")
	} else {
		ws("Rejects")
	}
}
