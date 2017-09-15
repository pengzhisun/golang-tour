/* ****************************************************************************
 * Tour: Nil interface values
 * Link: https://tour.golang.org/methods/13
 * ----------------------------------------------------------------------------
 * A nil interface value holds neither value nor concrete type.
 *
 * Calling a method on a nil interface is a run-time error because there is no
 * type inside the interface tuple to indicate which concrete method to call.
 * ***************************************************************************/

package main

import (
	"fmt"
)

// I interface
type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)
	i.M()
}
