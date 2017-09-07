/* ****************************************************************************
 * Tour: Exercise: Maps
 * Link: https://tour.golang.org/moretypes/23
 * ----------------------------------------------------------------------------
 * Implement WordCount. It should return a map of the counts of each “word”
 * in the string s. The wc.Test function runs a test suite against the provided
 * function and prints success or failure.
 *
 * You might find strings.Fields helpful.
 * ***************************************************************************/

package main

import "golang.org/x/tour/wc"
import "strings"

// WordCount func
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)

	for _, w := range words {
		result[w]++
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
