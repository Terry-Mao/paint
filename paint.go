package paint

import (
	"math"
	"paint/magickwand"
)

func Thumbnail(wand *magickwand.MagickWand, width, height uint) error {
    srcWidth := wand.Width()
    srcHeight := wand.Height()
	ratio := math.Min(float64(srcWidth)/float64(width),
		float64(srcHeight)/float64(height))
	scaleWidth := uint(float64(srcWidth) / ratio)
	scaleHeight := uint(float64(srcHeight) / ratio)

	// Resize
	if err := wand.Resize(scaleWidth, scaleHeight, magickwand.GaussianFilter,
		1.0); err != nil {
		return err
	}

	// Crop
	return wand.Crop(width, height, 0, 0)
}
