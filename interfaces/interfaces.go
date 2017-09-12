/* ****************************************************************************
 * Tour: Interfaces
 * Link: https://tour.golang.org/methods/9
 * ----------------------------------------------------------------------------
 * An interface type is defined as a set of method signatures.
 *
 * A value of interface type can hold any value that implements those methods.
 * ***************************************************************************/

package main

import (
	"fmt"
	"math"
)

// Abser interface
type Abser interface {
	Abs() float64
}

// MyFloat type
type MyFloat float64

// Abs method for MyFloat type
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex struct
type Vertex struct {
	X, Y float64
}

// Abs method for Vertex type
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}
