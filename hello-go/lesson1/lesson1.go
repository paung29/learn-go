package main

import "fmt"

func main() {
	var a string
	var b int 
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	greet("Aike")

	fmt.Println(adding(2,3))

	var person1 string
	person1 = "Deborah"

	fmt.Println(person1)
}

func greet(name string) {
	fmt.Println("Hello ", name)
}

func adding(a int, b int) int{
	return  a + b
}


