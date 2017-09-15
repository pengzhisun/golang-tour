/* ****************************************************************************
 * Exercise: rot13Reader
 * Link: https://tour.golang.org/methods/23
 * ----------------------------------------------------------------------------
 * A common pattern is an io.Reader that wraps another io.Reader, modifying the
 * stream in some way.
 *
 * For example, the gzip.NewReader function takes an io.Reader (a stream of
 * compressed data) and returns a *gzip.Reader that also implements io.Reader
 * (a stream of the decompressed data).
 *
 * Implement a rot13Reader (https://en.wikipedia.org/wiki/ROT13) that
 * implements io.Reader and reads from an  * io.Reader, modifying the stream by
 * applying the rot13 substitution cipher to all alphabetical characters.
 *
 * The rot13Reader type is provided for you. Make it an io.Reader by
 * implementing its Read method.
 * ***************************************************************************/

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Read method for rot13Reader struct
func (rot13 rot13Reader) Read(p []byte) (int, error) {
	n, err := rot13.r.Read(p)

	for i, v := range p[:n] {
		switch {
		case (v >= 'A' && v <= 'M') || (v >= 'a' && v <= 'm'):
			p[i] = p[i] + 13
		case (v >= 'N' && v <= 'Z') || (v >= 'n' && v <= 'z'):
			p[i] = p[i] - 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
