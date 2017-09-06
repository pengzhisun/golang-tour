/* ****************************************************************************
 * Tour: For continued
 * Link: https://tour.golang.org/flowcontrol/2
 * ----------------------------------------------------------------------------
 * The init and post statement are optional.
 * ***************************************************************************/

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
