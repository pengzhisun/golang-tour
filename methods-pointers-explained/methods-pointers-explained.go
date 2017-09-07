/* ****************************************************************************
 * Tour: Pointers and functions
 * Link: https://tour.golang.org/methods/5
 * ----------------------------------------------------------------------------
 * Here we see the Abs and Scale methods rewritten as functions.
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

// Scale func
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
