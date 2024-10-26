package util

import "github.com/gopxl/pixel/v2"

// helper function to convert a pixel.Vec to a string for direction, helpful for
// debugging and sprite key setting, returns 'zero' when the vector is zero
func Direction(v pixel.Vec) string {
	if v.X > 0 {
		if v.Y > 0 {
			return "right_up"
		} else if v.Y < 0 {
			return "right_down"
		} else {
			return "right"
		}
	} else if v.X < 0 {
		if v.Y > 0 {
			return "left_up"
		} else if v.Y < 0 {
			return "left_down"
		} else {
			return "left"
		}
	} else {
		if v.Y > 0 {
			return "up"
		} else if v.Y < 0 {
			return "down"
		}
	}

	return "zero"
}
