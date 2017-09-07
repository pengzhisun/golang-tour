/* ****************************************************************************
 * Tour: Methods are functions
 * Link: https://tour.golang.org/methods/2
 * ----------------------------------------------------------------------------
 * Remember: a method is just a function with a receiver argument.
 * Here's Abs written as a regular function with no change in functionality.
 * ***************************************************************************/

package main

import (
	"fmt"
	"math"
)

// Vertex struct
type Vertex struct {
	X, Y float64
}

// Abs func
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
