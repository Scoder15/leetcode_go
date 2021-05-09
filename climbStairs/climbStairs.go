package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	res := 0
	begin := 1
	end := 2
	for i := 3; i <= n; i++ {
		res = begin + end
		begin = end
		end = res

	}
	return res
}
