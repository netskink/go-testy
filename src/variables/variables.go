package main

// no commas in the list of imports
import (
	"fmt"
)

func main() { 
	// Every variable has an intialized zero value
	var  x int
	fmt.Println("x = ", x);
	var  y int = 3
	fmt.Println("y = ", y);
	// Shorthand 
	z := 33
	fmt.Println("z = ", z);
}
