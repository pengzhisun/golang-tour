/* ****************************************************************************
 * Tour: Numeric Constants
 * Link: https://tour.golang.org/basics/16
 * ----------------------------------------------------------------------------
 * Numeric constants are high-precision values.
 * An untyped constant takes the type needed by its context.
 * (An int can store at maximum a 64-bit integer, and sometimes less.)
 * ***************************************************************************/

package main

import "fmt"

const (
	// Big constants
	Big = 1 << 100
	// Small constants
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x + 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))

	// constant 1267650600228229401496703205376 overflows int
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
}
