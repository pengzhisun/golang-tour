/* ****************************************************************************
 * Exercise: Images
 * Link: https://tour.golang.org/methods/25
 * ----------------------------------------------------------------------------
 * Let's write another Image, this time it will return an implementation of
 * image.Image instead of a slice of data.
 *
 * Define your own Image type, implement the necessary methods, and call
 * pic.ShowImage.
 *
 * Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
 *
 * ColorModel should return color.RGBAModel.
 *
 * At should return a color; the value v in the last picture generator
 * corresponds to color.RGBA{v, v, 255, 255} in this one.
 *
 * Hint: https://codebeautify.org/base64-to-image-converter
 * ***************************************************************************/

package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image struct
type Image struct {
	w         int
	h         int
	generator func(int, int) uint8
}

// Bounds method for Image struct
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

// ColorModel method for Image struct
func (img Image) ColorModel() color.Model {
	return color.RGBA64Model
}

// At method for Image struct
func (img Image) At(x, y int) color.Color {
	v := img.generator(x, y)
	return color.RGBA{v, v, 255, 255}
}

func generator0(x, y int) uint8 {
	return 0
}

func generator1(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func generator2(x, y int) uint8 {
	return uint8(x * y)
}

func generator3(x, y int) uint8 {
	return uint8(x ^ y)
}

func main() {
	var width, height int = 256, 256

	// default image: images/images0.png https://git.io/v57kQ
	m0 := Image{width, height, generator0}
	fmt.Println("default image:")
	pic.ShowImage(m0)

	// (x+y)/2 image: images/images1.png https://git.io/v57Ie
	m1 := Image{width, height, generator1}
	fmt.Println("(x+y)/2 image:")
	pic.ShowImage(m1)

	// x*y image: images/images2.png https://git.io/v57Ik
	m2 := Image{width, height, generator2}
	fmt.Println("x*y image:")
	pic.ShowImage(m2)

	// x^y image: images/images3.png https://git.io/v57Is
	m3 := Image{width, height, generator3}
	fmt.Println("x^y image:")
	pic.ShowImage(m3)
}
