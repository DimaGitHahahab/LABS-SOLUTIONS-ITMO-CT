package main

import (
	"fmt"
)

func main() {
	var sum int
	var n, x, y, a0 int
	fmt.Scan(&n, &x, &y, &a0)
	a := make([]int, n)
	a[0] = a0
	for i := 1; i < n; i++ {
		a[i] = (x*a[i-1] + y) % 65536
		if a[i] < 0 {
			a[i] += 65536
		}
	}
	var m, z, t, b0 int
	fmt.Scan(&m, &z, &t, &b0)
	b := make([]int, 2*m)
	if m > 0 {
		b[0] = b0
		for i := 1; i < 2*m; i++ {
			b[i] = (z*b[i-1] + t) % 1073741824
			if b[i] < 0 {
				b[i] += 1073741824
			}
		}
		prefixSum := 0
		for i := 0; i < n; i++ {
			prefixSum += a[i]
			a[i] = prefixSum
		}
		for i := 0; i < m; i++ {
			l, r := b[2*i]%n, b[2*i+1]%n
			if l < 0 {
				l += n
			}
			if r < 0 {
				r += n
			}
			if l > r {
				l, r = r, l
			}
			if l == 0 {
				sum += a[r]
			} else {
				sum += a[r] - a[l-1]
			}
		}
	}
	fmt.Println(sum)
}
