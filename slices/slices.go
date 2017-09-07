/* ****************************************************************************
 * Tour: Slices
 * Link: https://tour.golang.org/moretypes/7
 * ----------------------------------------------------------------------------
 * An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
 * flexible view into the elements of an array. In practice, slices are much
 * more common than arrays.
 *
 * The type []T is a slice with elements of type T.
 *
 * This expression creates a slice of the first five elements of the array a:
 *   a[0:5]
 * ***************************************************************************/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[1:4]
	fmt.Println(s)
}
