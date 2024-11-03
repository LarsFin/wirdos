package util

import "github.com/gopxl/pixel/v2"

// Returns the minimum required translation to ensure b contains a, should a
// be larger than b for either or both dimensions the translation to anchor a
// on b's center will be returned for the respective dimension. If a is already
// within b, then a zero vector is returned as no translation is needed.
func ContainmentTranslation(a, b pixel.Rect) pixel.Vec {
	x, y := .0, .0

	if (a.W() > b.W()) {
		x = b.Center().X - a.Center().X
	} else if (a.Min.X < b.Min.X) {
		x = b.Min.X - a.Min.X
	} else if (a.Max.X > b.Max.X) {
		x = b.Max.X - a.Max.X
	}

	if (a.H() > b.H()) {
		y = b.Center().Y - a.Center().Y
	} else if (a.Min.Y < b.Min.Y) {
		y = b.Min.Y - a.Min.Y
	} else if (a.Max.Y > b.Max.Y) {
		y = b.Max.Y - a.Max.Y
	}

	return pixel.V(x, y)
}
