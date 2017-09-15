/* ****************************************************************************
 * Exercise: Errors
 * Link: https://tour.golang.org/methods/20
 * ----------------------------------------------------------------------------
 * Sqrt should return a non-nil error value when given a negative number, as it
 * doesn't support complex numbers.
 *
 * Create a new type
 *   type ErrNegativeSqrt float64
 *
 * and make it an error by giving it a
 *   func (e ErrNegativeSqrt) Error() string
 *
 * method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative
 * number: -2".
 *
 * Note: a call to fmt.Sprint(e) inside the Error method will send the program
 * into an infinite loop. You can avoid this by converting e first:
 * fmt.Sprint(float64(e)).
 *
 * Change your Sqrt function to return an ErrNegativeSqrt value when given a
 * negative number.
 * ***************************************************************************/

package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt type
type ErrNegativeSqrt float64

// Error method for ErrNegativeSqrt type
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

// SqrtZ returns the square root of x using Newton's method
func SqrtZ(x, z float64) float64 {
	return z - (math.Pow(z, 2)-x)/(2*z)
}

// Sqrt returns the square root of x
func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	i := 1
	changing := true
	for changing {
		z1 := SqrtZ(x, z)
		// fmt.Printf("loop: %v, value: %g\n", i, z1)

		changing = math.Abs(z1-z) > 0.000000000000001
		z = z1
		i++
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
