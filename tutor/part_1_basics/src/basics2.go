package main
// programs start running in package main.

import (
    "fmt"
)

// go's return values may be named.  if so they are treated as variables at the
// top of the function.
//
// A return statement without arguments returns the named return variables.
// This is known as a naked return

func split (sum int)(x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}


func main() {
    fmt.Println("split(27) is")
    fmt.Println(split(27))
}


