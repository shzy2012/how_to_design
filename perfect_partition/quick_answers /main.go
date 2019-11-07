package main

import (
	"fmt"
	"math"
)

/* 完美分区解法
http://mathworld.wolfram.com/PerfectPartition.html
http://oeis.org/A002033/list
通过映射,将Perfect Partition问题映射为A002033表
*/
/*
A002033 as a simple table
n		a(n)
0		1
1		1
2		1
3		2
4		1
5		3
6		1
7		4
8		2
9		3
10		1
11		8
12		1
13		3
14		3
15		8
16		1
17		8
18		1
19		8
20		3
*/

func modp(a, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}

func proc(n int) int {
	if n <= 2 {
		return 1
	}

	a := 0
	for i := 0; i < n-1; i++ {
		if modp(n+1, i+1) == 0 {
			tmp := a
			a = a + proc(i)
			fmt.Printf("f(%v)=%v+%v+%v\n", n, a, i, tmp)
		}
	}

	return a
}

func main() {
	for n := 5; n <= 6; n++ {
		out := proc(n)
		fmt.Printf("f(%v)=%v\n", n, out)
	}
}
