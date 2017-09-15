/* ****************************************************************************
 * Tour: Interface values with nil underlying values
 * Link: https://tour.golang.org/methods/12
 * ----------------------------------------------------------------------------
 * If the concrete value inside the interface itself is nil, the method will be
 * called with a nil receiver.
 *
 * In some languages this would trigger a null pointer exception, but in Go it
 * is common to write methods that gracefully handle being called with a nil
 * receiver (as with the method M in this example.)
 *
 * Note that an interface value that holds a nil concrete value is itself
 * non-nil.
 * ***************************************************************************/

package main

import (
	"fmt"
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
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
