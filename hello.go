package main

import "fmt"

type Address struct{
	country string
	city string
}

type User struct {
	name string
	age int
	location Address
}

func main() {
	loc:=Address{"Uzbekistan", "Kokand"}
	bob:=User {"Xurshid", 17, loc}
	fmt.Println("Name", bob.name)
	fmt.Println("Age", bob.age)
	fmt.Println("Loc", bob.location)
}
