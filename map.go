package main

import "fmt"

func main() {
	var a = make(map[string]string)
	var b map[string]string
	a["brand"] = "ford"
	a["model"] = "mustang"
	a["year"] = "1964"

	b["brand"] = "ford"
	b["model"] = "mustang"
	b["year"] = "1964"

	fmt.Println(a["brand"])

	fmt.Println(b["brand"])

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
