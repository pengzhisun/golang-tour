/* ****************************************************************************
 * Tour: if-and-else
 * Link: https://tour.golang.org/
 * ----------------------------------------------------------------------------
 * Variables declared inside an if short statement are also available inside
 * any of the else blocks.
 *
 * Both calls to pow are executed and return before the call to fmt.Println
 * in main begins.
 * ***************************************************************************/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// golint warning: if block ends with a return statement, so drop this
		// else and outdent its block (move short variable declaration to its
		// own line if necessary)
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func powLintFree(x, n, lim float64) float64 {
	v := math.Pow(x, n)

	if v < lim {
		return v
	}
	fmt.Printf("%g >= %g\n", v, lim)
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
