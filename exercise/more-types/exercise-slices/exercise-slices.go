/* ****************************************************************************
 * Tour: Exercise: Slices
 * Link: https://tour.golang.org/moretypes/18
 * ----------------------------------------------------------------------------
 * Implement Pic. It should return a slice of length dy, each element of which
 * is a slice of dx 8-bit unsigned integers. When you run the program, it will
 * display your picture, interpreting the integers as grayscale (well,
 * bluescale) values.
 *
 * The choice of image is up to you. Interesting functions include (x+y)/2,
 * x*y, and x^y.
 *
 * You need to use a loop to allocate each []uint8 inside the [][]uint8.
 *
 * Use uint8(intValue) to convert between types.
 * ***************************************************************************/

package main

import "golang.org/x/tour/pic"

// Pic func
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for y := range s {
		s[y] = make([]uint8, dx)
		for x := range s[y] {
			// default image: https://raw.githubusercontent.com/pengzhisun/golang-tour/master/exercise/more-types/exercise-slices/images/slices0.png

			// (x+y)/2 image: https://raw.githubusercontent.com/pengzhisun/golang-tour/master/exercise/more-types/exercise-slices/images/slices1.png
			// s[y][x] = uint8((x + y) / 2)

			// x*y image: https://raw.githubusercontent.com/pengzhisun/golang-tour/master/exercise/more-types/exercise-slices/images/slices2.png
			// s[y][x] = uint8(x * y)

			// x^y image: https://raw.githubusercontent.com/pengzhisun/golang-tour/master/exercise/more-types/exercise-slices/images/slices3.png
			s[y][x] = uint8(x ^ y)
		}
	}

	return s
}

func main() {
	pic.Show(Pic)
}
