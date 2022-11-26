// package main

// import (
// 	"fmt"
// )

// func isprime(num int) bool {
// 	for i := 2; i < num/2; i++ {
// 		fmt.Println(i)
// 		if num%i == 0 {
// 			return false
// 		}

// 	}
// 	return true

// }

// func main() {
// a := [...]int{10, 20, 30, 40, 50}
// fmt.Println("the array is \n", a)
// a[0] = 10
// a[1] = 20
// fmt.Println("the value of array: \n", a)
// a[1] = 0
// fmt.Println("the value of array after changing: \n", a)

// b := &a

// a[4] = 60

// fmt.Println("a: ", a)
// fmt.Println("b: ", b)

// for i, v := range a {
// 	fmt.Printf("the index is i:%d and the value is :%d \n", i, v)
// 	fmt.Println("\nthe index is i:", i, "the value of v: ", v)
// }

// fmt.Println(len(a))

// fmt.Println("a==b", a == *b)

// grades := [3]int{0: 10, 1: 20, 2: 30}
// fmt.Println(grades)
// fmt.Printf("%#v", grades)

// fmt.Println(isprime(7))
// var sl []int

// fmt.Println("declaring a slice of int: ", sl == nil)
// fmt.Printf("cities: %#v\n", sl)

// sl[0] = 10
// fmt.Println(sl)

// s := []string{"moin", "suhas", "nandu"}
// fmt.Println(s)
// fmt.Printf("%#v", s)
// fmt.Println()

// for _, val := range s {
// 	fmt.Printf(" the value is %s ", val)
// }
// fmt.Println()
// for i := 0; i < len(s); i++ {
// 	fmt.Println(s[i])

// }

// a := []int{1, 2, 3, 4, 5}
// a1 := a[1:3]
// fmt.Printf("the values of a is %d and the value of a1 is %d: ", a, a1)
// a1[0] = 10
// fmt.Println()
// fmt.Printf("the values of a is %d and the value of a1 is %d:\n ", a, a1)
// fmt.Printf("the type of a1 is %T", a1)
// a = append(a[:5], 100)
// fmt.Println(a)
// a = append(a[:7], 20)
// fmt.Println(a)
// fmt.Println(len(a))
// s := "moin ali khan"
// a := 0
// for _, val := range s {
// 	a += 1
// 	fmt.Printf("%c", val)

// }
// fmt.Println()
// b := utf8.RuneCountInString(s)
// fmt.Println(b)

// for i := 0; i < len(s); i++ {
// 	fmt.Printf("%c", s[i])

// }
// fmt.Println()
// fmt.Println(strings.Repeat("#", 20))

// 	type book struct{
// 		title string
// 		author string
// 		lunch date int
// 	}

// 	type boo1 struct{
// 		title string
// 		author string
// 		lunch date int
// 	}

// 	latestbook:= book{"wisdom","danish",1997}
// 	lastbook:=book1{"storybook","jenny",2000}

// 	fmt.Println("the latest book of danish is: ",book)
// }

// type book struct {
// 	title, author string
// 	year          int
// }

// latest := book{"wisdom", "danish", 2000}
// fmt.Println(latest)

// latest.title = "wisdom1"
// fmt.Println(latest)

// fmt.Println(latest.author)

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// type LuckyNumber struct {
// 	number int
// }

// type Person struct {
// 	lucky_numbers []LuckyNumber
// }

// func main() {
// 	count_of_lucky_nums := 10
// 	// START OF SECTION I WANT TO OPTIMIZE
// 	var tmp []LuckyNumber
// 	for i := 0; i < count_of_lucky_nums; i++ {
// 		tmp = append(tmp, LuckyNumber{rand.Intn(100)})
// 	}
// 	a := Person{tmp}
// 	// END OF SECTION I WANT TO OPTIMIZE
// 	fmt.Println(a)
// }

package main

import "fmt"

func main() {

	dana := struct {
		firstname, lastname string
		age                 int
	}{
		firstname: "dana",
		lastname:  "n",
		age:       24,
	}
	fmt.Printf("%#v/n", dana)
	fmt.Println(dana)

	type book struct {
		string
		float64
		int
	}
	b := book{"hello", 10.5, 6}
	fmt.Println(b)

	type contact struct {
		mailid string
		phnum  int
	}
	type Employee struct {
		name        string
		empid       int
		contactInfo contact
	}
	moin := Employee{
		name:  "moin ali khan",
		empid: 2974,
		contactInfo: contact{
			mailid: "moinalikhan9.ma@gmail.com",
			phnum:  9590726059,
		},
	}
	fmt.Printf("%#v\n", moin)

	fmt.Println(moin.name)

}
