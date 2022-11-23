package main

import "fmt"

const seconds_in_hour = 3600

func main() {
	fmt.Println("hello world!")
	distance := 68.8
	fmt.Printf("the distance in miles is %f \n", distance*0.621)

	var age = 30
	fmt.Println("the age is", age)

	var name = "moin"
	fmt.Println("the name of the candidate is", name)
	// _ = name

	s := "moin is learning golang from udemy platform"
	fmt.Println(s)

	car, cost := "audi", 50000
	fmt.Println(car, cost)

	var (
		salary    float64
		firstname string
		gender    bool
	)

	fmt.Println(salary, firstname, gender)

	var a, b, c int
	println(a, b, c)

	var i, j int

	i, j = 5, 8
	fmt.Printf("after swapping %d,%d: ", i, j)
	fmt.Println(i, j)
	j, i = i, j
	fmt.Printf("before swapping %d,%d:", i, j)
	fmt.Println(i, j)
	_, _ = i, j

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		north = iota
		south
		east
		west
	)
	fmt.Println(north, south, east, west)
}
