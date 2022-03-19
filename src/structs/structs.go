package main

// no commas in the list of imports
import (
	"fmt"
)

type person struct {
	name string
	age int
}

// struct demo
func main() { 

	p := person{name:"john", age:55}

	fmt.Println(p)
	fmt.Println(p.name)


}


