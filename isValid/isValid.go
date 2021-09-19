package main

import "fmt"

func main() {

	fmt.Println(isValid("([}}])"))

}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0)
	m := make(map[byte]byte, 0)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			//fmt.Println(stack)
		} else if ((s[i] == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((s[i] == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((s[i] == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			//fmt.Println("stack:", len(stack))
			stack = stack[:len(stack)-1]

		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false

}
