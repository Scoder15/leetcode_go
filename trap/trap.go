package main

import(
	"fmt"
	"strconv"
)

func main() {
	num2 := "0"
	num1 := "0"
	fmt.Print(addStrings(num1,num2))
}

func addStrings(num1 string ,num2 string) string{
	l1 := len(num1) -1 
	l2 := len(num2) -1 
	temp := 0
	res := ""
	for l1 >= 0 && l2 >= 0{
		a,_ := strconv.Atoi(string(num1[l1]))
		b,_ := strconv.Atoi(string(num2[l2]))
		val := (a + b + temp) % 10  
		temp = (a + b + temp) / 10 
		valStr := strconv.Itoa(val)
		res = res + valStr
		l1--
		l2--
	}
	fmt.Println(l1,l2)
	if l1 == l2 && temp > 0{
		res = res + strconv.Itoa(temp)
		return reStr(res)
	}
	var p *string
	var l int
	
	if l1 > l2{
		p,l = &num1,l1
	}else if(l2 > l1){
		p,l = &num2,l2
	}

	isFirst := false
	for l >= 0 && p != nil{
		val,_ := strconv.Atoi(string((*p)[l]))
		if !isFirst{
			val = val + temp
			isFirst = true
		}
		valStr := strconv.Itoa(val)
		res = res + valStr
		l--
	}
	return reStr(res)
}

func reStr(s string) string{
	r := []rune(s)
	for i,j := 0,len(r)-1;i<j;{
		r[i],r[j] = r[j],r[i]
		i++
		j--
	}
	return string(r)
}