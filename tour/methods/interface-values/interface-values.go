/* ****************************************************************************
 * Tour: Interface values
 * Link: https://tour.golang.org/methods/11
 * ----------------------------------------------------------------------------
 * Under the covers, interface values can be thought of as a tuple of a value
 * and a concrete type:
 *   (value, type)
 *
 * An interface value holds a value of a specific underlying concrete type.
 *
 * Calling a method on an interface value executes the method of the same name
 * on its underlying type.
 * ***************************************************************************/

package main

import (
	"fmt"
	"math"
)

// I interface
type I interface {
	M()
}

// T struct
type T struct {
	S string
}

// M method for T struct
func (t *T) M() {
	fmt.Println(t.S)
}

// F type
type F float64

// M method for F type
func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}
