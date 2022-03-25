package main

// no commas in the list of imports
import (
	"fmt"
)

// arrays hold all the same type of variable and are unmutable
func main() { 

	// variable array of five ints
	var a [5]int
	a[2] = 7
    	fmt.Println("a is ", a);

	// Init array
	b := [5]int{5, 4, 3, 2, 1}
    	fmt.Println("b is ", b);


}
