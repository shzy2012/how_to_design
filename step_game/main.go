package main

import "fmt"

//1. 归纳推演
//公式 f(n) = f(n-1) + f(n-2)

//2. 递归实现
func rc(n int) int {
	if n <= 2 {
		return n
	}

	return rc(n-1) + rc(n-2)
}

//3. 动态规划实现
var steps []int

func dp(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		steps[n-1] = n
		return steps[n-1]
	}

	if steps[n-1] > 0 {
		return steps[n-1]
	}

	steps[n-1] = dp(n-1) + dp(n-2)
	return steps[n-1]
}

func main() {
	n := 10
	steps = make([]int, n)
	stepRC := rc(n) //递归
	stepDP := dp(n) //动态规划
	fmt.Printf("共有走法: %v vs %v", stepRC, stepDP)
}
