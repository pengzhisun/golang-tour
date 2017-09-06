/* ****************************************************************************
 * Exercise: Loops and Functions
 * Link: https://tour.golang.org/flowcontrol/8
 * ----------------------------------------------------------------------------
 * As a simple way to play with functions and loops, implement the square root
 * function using Newton's method.
 *
 * In this case, Newton's method is to approximate Sqrt(x) by picking a
 * starting point z and then repeating:
 *   Z_(n+1) = Z_n - (Z_n^2 - X)/2*Z_n
 *
 * To begin with, just repeat that calculation 10 times and see how close you
 * get to the answer for various values (1, 2, 3, ...).
 *
 * Next, change the loop condition to stop once the value has stopped changing
 * (or only changes by a very small delta). See if that's more or fewer
 * iterations. How close are you to the math.Sqrt?
 *
 * Hint: to declare and initialize a floating point value, give it floating
 * point syntax or use a conversion:
 *   z := float64(1)
 *   z := 1.0
 * ***************************************************************************/

package main

import (
	"fmt"
	"math"
)

// SqrtZ returns the square root of x using Newton's method
func SqrtZ(x, z float64) float64 {
	return z - (math.Pow(z, 2)-x)/(2*z)
}

// Sqrt returns the square root of x
func Sqrt(x float64) float64 {
	z := 1.0
	i := 1
	changing := true
	for changing {
		z1 := SqrtZ(x, z)
		fmt.Printf("loop: %v, value: %g\n", i, z1)

		changing = math.Abs(z1-z) > 0.000000000000001
		z = z1
		i++
	}
	return z
}

func main() {
	fmt.Printf("Sqrt(2) result: %g\n", Sqrt(2))
	fmt.Printf("math.Sqrt(2) result: %g\n", math.Sqrt(2))
}
