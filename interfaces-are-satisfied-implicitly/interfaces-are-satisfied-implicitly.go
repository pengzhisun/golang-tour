/* ****************************************************************************
 * Tour: Interfaces are implemented implicitly
 * Link: https://tour.golang.org/methods/10
 * ----------------------------------------------------------------------------
 * A type implements an interface by implementing its methods. There is no
 * explicit declaration of intent, no "implements" keyword.
 *
 * Implicit interfaces decouple the definition of an interface from its
 * implementation, which could then appear in any package without
 * prearrangement.
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
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
