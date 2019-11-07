### perfect partition

<div style="text-align: right"> 周毅 2019-10-31 </div>

### Introduction(背景故事)

> A perfect partition is a partition of a number n whose elements uniquely generate any number 1, 2, ..., n. {1,1,...,1_()_(n)} is always a perfect partition of n, and every perfect partition must contain a 1.

The following table gives the first several perfect partitions for small n.
```text
n	a(n)	perfect partitions
1	 1	    {1}
2	 1	    {1,1}
3	 2	    {2,1}, {1,1,1}
4	 1	    {1,1,1,1}
5	 3	    {3,1,1}, {2,2,1}, {1,1,1,1,1}
6	 1	    {1,1,1,1,1,1}
```
---

### Requirements (需求列表、test cases)

 1. 给定N，打印a(n)和全部perfect partitions
 2. 输入输出如下

```text
input N:3
answer:
a(3) = 2
p#1 = {1,2}
p#2 = {1,1,1}
```
---

### Design flow(证明)
```text
1. 获取输入参数，如果没有默认1
2. 寻找输入参数的完美解并打印输出
    2.1 利用递归寻找完美解
        2.1.1 n=0或n=1时直接返回结果
        2.1.2 n > 1时利用动规求解
```

<img src="https://raw.githubusercontent.com/shzy2012/static/master/pp_solution.png?raw=true" width="600" height="350">

---


### Alternatives & Analysis
决策：实现采用动态规划
```text
  逻辑清晰、代码简单，容易实现，程序运行效率高
```

---

### Risk management
1. 如果缓存可能过大可能导致内存不够
