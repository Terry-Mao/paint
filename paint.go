// Copyright 2013 Terry.Mao. All rights reserved.

// Package paint providers most used image function.
package paint

import (
	"github.com/Terry-Mao/paint/wand"
)

// Converts the current image into a thumbnail of the specified width and 
// height. 
func Thumbnail(w *wand.MagickWand, width, height uint) error {
	// Resize
	return w.ResizeImage(width, height, wand.GaussianFilter, 1.0)
}
