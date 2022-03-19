package main

// no commas in the list of imports
import (
	"fmt"
)

type person struct {
	name string
	age int
}

// pointers demo
func main() { 

	i := 7
	fmt.Println(i)
	fmt.Println(&i)

	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}
