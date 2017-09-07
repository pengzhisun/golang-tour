/* ****************************************************************************
 * Tour: Structs
 * Link: https://tour.golang.org/moretypes/2
 * ----------------------------------------------------------------------------
 * A struct is a collection of fields.
 * (And a type declaration does what you'd expect.)
 * ***************************************************************************/

package main

import "fmt"

// Vertex struct
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
