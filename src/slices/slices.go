package main

// no commas in the list of imports
import (
	"fmt"
)

// slices are a mutable form of arrays
func main() { 

	// variable array of five ints
	// similar syntax.  Just remove the length specifier
	var a []int
	// now we can use append
	a = append(a, 3)
    	fmt.Println("a is ", a);

	// Init array
	b := []int{5, 4, 3, 2, 1}
	b[2] = 7
	b = append(b, 0)
    	fmt.Println("b is ", b);


}
