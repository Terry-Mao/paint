// Copyright 2013 Terry.Mao. All rights reserved.

// Package paint providers most used image function.
package paint

import (
	"github.com/Terry-Mao/paint/magickwand"
)

// Converts the current image into a thumbnail of the specified width and 
// height. 
func Thumbnail(wand *magickwand.MagickWand, width, height uint) error {
	// Resize
	return wand.ResizeImage(width, height, magickwand.GaussianFilter, 1.0)
}
