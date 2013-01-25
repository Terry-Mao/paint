package paint

import (
	"github.com/Terry-Mao/paint/magickwand"
)

func Thumbnail(wand *magickwand.MagickWand, width, height uint) error {
	// Resize
	return wand.Resize(width, height, magickwand.GaussianFilter, 1.0)
}
