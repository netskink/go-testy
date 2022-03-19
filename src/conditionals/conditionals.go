package main

// no commas in the list of imports
import (
	"fmt"
)

func main() { 
	x := 5
	if x > 5 {
	    fmt.Println("x greater than 5 ");
	} else if x < 5 {
	    fmt.Println("x less than 5 ");
	} else {
	    fmt.Println("x is 5 ");
	}
}
