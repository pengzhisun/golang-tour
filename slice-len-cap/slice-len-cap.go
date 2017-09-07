/* ****************************************************************************
 * Tour: Slice length and capacity
 * Link: https://tour.golang.org/moretypes/11
 * ----------------------------------------------------------------------------
 * A slice has both a length and a capacity.
 *
 * The length of a slice is the number of elements it contains.
 *
 * The capacity of a slice is the number of elements in the underlying array,
 * counting from the first element in the slice.
 *
 * The length and capacity of a slice s can be obtained using the expressions
 * len(s) and cap(s).
 *
 * You can extend a slice's length by re-slicing it, provided it has sufficient
 * capacity.
 * ***************************************************************************/

package main

import "fmt"

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlices(s)

	s = s[:0]
	printSlices(s)

	s = s[:4]
	printSlices(s)

	s = s[2:]
	printSlices(s)
}
