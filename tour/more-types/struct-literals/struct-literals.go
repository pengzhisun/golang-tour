/* ****************************************************************************
 * Tour: Struct Literals
 * Link: https://tour.golang.org/moretypes/5
 * ----------------------------------------------------------------------------
 * A struct literal denotes a newly allocated struct value by listing the values
 * of its fields.
 *
 * You can list just a subset of fields by using the Name: syntax. (And the
 * order of named fields is irrelevant.)
 *
 * The special prefix & returns a pointer to the struct value.
 * ***************************************************************************/

package main

import "fmt"

// Vertex struct
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
