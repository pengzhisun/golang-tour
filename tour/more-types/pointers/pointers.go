/* ****************************************************************************
 * Tour: Pointers
 * Link: https://tour.golang.org/moretypes/1
 * ----------------------------------------------------------------------------
 * Go has pointers. A pointer holds the memory address of a value.
 *
 * The type *T is a pointer to a T value. Its zero value is nil.
 *   var p *int
 *
 * The & operator generates a pointer to its operand.
 *   i := 42
 *   p = &i
 *
 * The * operator denotes the pointer's underlying value.
 *   fmt.Println(*p) // read i through the pointer p
 *   *p = 21         // set i through the pointer p
 *
 * This is known as "dereferencing" or "indirecting".
 *
 * Unlike C, Go has no pointer arithmetic.
 * ***************************************************************************/

package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Printf("read i though the pointer: %v\n", *p)
	*p = 21
	fmt.Printf("see the new value of i: %v\n", i)

	p = &j
	*p = *p / 37
	fmt.Printf("see the new value of j: %v\n", j)
}
