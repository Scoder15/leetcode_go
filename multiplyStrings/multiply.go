package main

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/multiply-strings/
func main() {
	a := "0"
	b := "000"
	fmt.Println(multiply(a, b))

}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := "0"
	resStrs := make([]string, 0)
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	count := 0
	fmt.Println("multiply running")
	for j := l2; j >= 0; j-- {
		// fmt.Println("j:", j)
		tempRes := ""
		carry := 0
		for i := l1; i >= 0; i-- {

			// fmt.Println("i:", i)
			x, _ := strconv.Atoi(string(num1[i]))
			y, _ := strconv.Atoi(string(num2[j]))
			temp := x*y + carry
			carry = temp / 10
			val := temp % 10
			tempRes += strconv.Itoa(val)
		}
		if carry > 0 {
			tempRes += strconv.Itoa(carry)
		}
		reRes := reStr(tempRes)
		k := count
		for k > 0 {
			reRes += "0"
			k--
		}
		if tempRes != "" {
			count++
		}
		resStrs = append(resStrs, reRes)
	}
	for i := 0; i < len(resStrs); i++ {
		res = addStrings(res, resStrs[i])
	}
	return res
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
	if carry == 1 {
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
