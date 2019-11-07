// 题目：Prefect Partition
// 作者：周毅 2019-11-07
// 背景：对整数n进行一种划分，使其能够通过加法产生1,2,......,n，并且产生方式只有一种
//
// 用动态规划来求解（满足无后效性）
// f(n) = n-k U f(k)  条件：n-k = max(f(k)) or n-k > sum(f(k)), n-k <= sum(f(k))+1
// 边界
// f(0) = {0}
// f(1) = {1}

package main

import (
	"fmt"
	"math"
	"strings"
)

//缓存，记录k的解
var cache map[int][][]int

// Perfect Partition:
//
// 1. 获取输入参数，如果没有默认1
// 2. 寻找输入参数的完美解并打印输出
//     2.1 利用递归寻找完美解
//         2.1.1 n=0或n=1时直接返回结果
//         2.1.2 n > 1时利用动规求解
func pp(n int) [][]int {

	if aa, err := cache[n]; err {
		return aa
	}

	res := [][]int{}
	if n == 0 {
		res = append(res, []int{0})
	} else if n == 1 {
		res = append(res, []int{1})
	} else {

		for k := int(n / 2); k < n; k++ {
			solulist := pp(k)
			cache[k] = solulist
			for _, sublist := range solulist {
				//判断条件2
				if n-k == Max(sublist) || n-k > Sum(sublist) {
					//f(n) = (n-k) U f(k)
					sublist = append([]int{n - k}, sublist...)
					res = append(res, sublist)
				}
			}
		}
	}

	return res
}

//Sum Get sum of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum = sum + v
	}

	return sum
}

//Max Get max value from numbers
func Max(numbers []int) int {
	max := math.MinInt64
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}

func main() {

	cache = make(map[int][][]int, 0)
	n := 10
	result := pp(n)

	fmt.Printf("Input N:%v\n", n)
	fmt.Println("Answer:")
	fmt.Printf("a(%v) = %v\n", n, len(result))
	for i, item := range result {
		fmt.Printf("p#%v = {%v}\n", i+1, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(item)), ","), "[]"))
	}
}
