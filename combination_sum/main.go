package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	sulution := []int{}
	cs(candidates, sulution, target, &res)
	return res
}

func cs(candidates, solution []int, target int, result *[][]int) {

	if target == 0 {
		*result = append(*result, solution)
	}

	if len(candidates) <= 0 || target < candidates[0] {
		// target < candidates[0] 因为candidates是排序好的
		return
	}

	for i, ele := range candidates {

		// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
		// 避免多处同时对底层数组进行修改，产生错误的答案
		solution = solution[:len(solution):len(solution)]
		cs(candidates[i:], append(solution, ele), target-ele, result)
	}
}

func main() {

	candidates := []int{2, 3, 5}
	target :=8
	list := combinationSum(candidates, target)

	for _, arr := range list {
		fmt.Printf("%v\n", arr)
	}
}
