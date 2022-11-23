package main

import "fmt"

func Contains(x []string, s string) bool {
	for _, v := range x {
		if v == s {
			return true
		}
	}
	return false
}

func isValid(s string) bool {
	var a []rune
	var p = []string{"[]", "()", "{}"}
	var t string
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			a = append(a, v)
		} else {
			if len(a) > 0 {
				t = fmt.Sprintf("%c%c", a[len(a)-1], v)
				if Contains(p, t) {
					a = a[:len(a)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	if len(a) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Print(isValid("()"))
}
