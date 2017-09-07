/* ****************************************************************************
 * Tour: Map literals
 * Link: https://tour.golang.org/moretypes/20
 * ----------------------------------------------------------------------------
 * Map literals are like struct literals, but the keys are required.
 * ***************************************************************************/

package main

import "fmt"

// Vertex struct
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433,
		-74.39967,
	},
	"Google": Vertex{
		37.42202,
		-122.08408,
	},
}

func main() {
	fmt.Println(m)
}
