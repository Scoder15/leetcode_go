package main

import (
	"fmt"
	"strconv"
)
//https://leetcode-cn.com/problems/add-strings/

func main() {
	a := "999"
	b := "9"
	fmt.Println(a[1]-'0', b[0]-'0')
	fmt.Println(addStrings(a, b))
}

func addStrings(num1 string, num2 string) string {
	res := ""
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	carry := 0
	for l1 >= 0 || l2 >= 0 {
		x := 0
		y := 0
		if l1 >= 0 {
			x, _ = strconv.Atoi(string(num1[l1]))

		}

		if l2 >= 0 {

			y, _ = strconv.Atoi(string(num2[l2]))

		}

		temp := x + y + carry
		carry = temp / 10
		val := temp % 10
		res += strconv.Itoa(val)
		l1--
		l2--
	}
	if carry == 1{
		res += strconv.Itoa(carry)
	}
	return reStr(res)

}

func reStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
