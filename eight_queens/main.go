package main

import (
	"fmt"
)

var n int
var queen []int
var count int

/*检查相应位置是否可以放置皇后*/
func checkPos(row, col int) int {
	for k := 0; k < row; k++ {
		if queen[k] == col || queen[k]-col == k-row || queen[k]-col == row-k {
			return 0
		}
	}

	return 1
}

/*打印皇后的位置*/
func print(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if queen[i] == j {
				fmt.Print(" Q ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}

	fmt.Println("---------------------------------------------------------------------")
}

//递归遍历所有的情况
func find(row int) {
	//当row == n时表明已放置了n个皇后，递归结束，记录一个解
	if row == n {
		count = count + 1
		//8.打印放置好的棋盘
		print(n)
	} else {
		//4.逐一查找每一行能放皇后的位置。
		for col := 0; col < n; col++ {
			//5.检查相应位置是否可以放置皇后
			if checkPos(row, col) == 1 {
				//6.放置皇后
				queen[row] = col
				//7.找到第 i 行的解的前提是前 i-1 行都放好了皇后，找到第 N 行的时候，递归结束
				find(row + 1)
			}
		}
	}
}

func main() {

	count = 0
	//1.定义皇后个数
	n = 8
	//2.初始化记录皇后的数组
	queen = make([]int, n)
	//3.开始执行从0放置皇后
	find(0)
	//9.打印总数
	fmt.Println("total:", count)
}
