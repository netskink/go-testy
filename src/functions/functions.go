package main

// no commas in the list of imports
import (
	"fmt"
	"errors"
	"math"
)

func mysqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
// the return type of the function is after params list
func main() { 


	// a simple function with args
	x := 3
	y := 2

	fmt.Println("sum = ", sum(x,y) )




	// can also have functions which return multiple values
	var result float64
	var err  error
	result, err = mysqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}


	// This uses the error
	result, err = mysqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sum(x int, y int) int {
	return x + y
}

