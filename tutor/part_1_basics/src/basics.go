package main
// programs start running in package main.

import (
    "fmt"
    "math/rand"
    "math"
)

// by convention the package name is the same as the last element of
// the import path.  For insance "math/rand" package
// comprises files that begin with the statement
// package rand.

// notice that function parameters follow
// the variable name. the function return 
// type comes after the final parameter parenthesis.
func add(x int, y int) int {
    return x + y
}

// consecutive parameters of the same type can
// be short cutted like so
func sub(x,y int) int {
    return x - y
}

// a function can return more than one result.
// This function returns two strings
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    rand.Seed(331967)
    fmt.Println("My favorite number is", rand.Intn(10))
    // A name is exported if it begins with a captial letter
    // Any non capitalized letter names are not exported and
    // are not available outside the packages
    fmt.Println("My most favorite number is", math.Pi)
    fmt.Println("one plus two is", add(1,2))
    fmt.Println("one minus two is", sub(1,2))
    a, b := swap("one", "two")
    fmt.Println("swap(one two) is")
    fmt.Println(a,b)
}


