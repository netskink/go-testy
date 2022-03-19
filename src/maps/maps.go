package main

// no commas in the list of imports
import (
	"fmt"
)

// maps are a hashmap
func main() { 

	// first type is for key
	// second type is for value
	// use builtin make() function
	vertices := make(map[string]int)
	// now we can assign key value pairs
	// Need to use double quotes 
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["rectangle"] = 4;
    	fmt.Println("vertices is ", vertices);
	// NOTE: spaces are automatically added
    	fmt.Println("triangles have", vertices["triangle"], "vertices");
	// can delete keys
	delete(vertices,"rectangle")
    	fmt.Println("vertices is ", vertices);



}
