package main

import "fmt"

name := "Brad"

func main() {
//	 using var
//	var name string = "Brad"
	var age int = 37
	const isCool = true

//Shorthand
//	name := "Brad"
//	size := 1.3
//	email := "brad@gmail.com"
//
	name,email := "Brad","brad@gmail.com"

	fmt.Println(name,age,isCool)
	fmt.Printf("%T\n",email)
}
