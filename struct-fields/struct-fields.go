/* ****************************************************************************
 * Tour: Struct Fields
 * Link: https://tour.golang.org/moretypes/3
 * ----------------------------------------------------------------------------
 * Struct fields are accessed using a dot.
 * ***************************************************************************/

package main

import "fmt"

// Vertex struct
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
