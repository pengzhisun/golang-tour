/* ****************************************************************************
 * Exercise: Fibonacci closure
 * Link: https://tour.golang.org/moretypes/26
 * ----------------------------------------------------------------------------
 * Let's have some fun with functions.
 *
 * Implement a fibonacci function that returns a function (a closure) that
 * returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
 * ***************************************************************************/

package main

import (
	"fmt"
)

func fibonacci() func() int {
	s := []int{0, 1}
	i := 0

	return func() int {
		var result int

		if i < 2 {
			result = s[i]
		} else {
			result = s[0] + s[1]
			s[0] = s[1]
			s[1] = result
		}

		i++
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
