package main

// no commas in the list of imports
import (
	"fmt"
)

// only one type of loop in go, but can be used
// as while loops
func main() { 

	// for loop
	for i := 0; i < 5; i++ {
	    fmt.Println("i is ", i);
	}

	// while loop
	j := 0
	for j < 5 {
	    fmt.Println("j is ", j);
	    j++
	}

	/// index over slice
	x := []string{"a", "b", "c"}
	
	for index, value := range x {
		fmt.Println("index", index, "value", value)
	}


	/// index over maps
	y := make(map[string]string)
	y["a"] = "alpha"
	y["b"] = "beta"
	y["c"] = "charlie"
	
	for key, value := range x {
		fmt.Println("key", key, "value", value)
	}




}
