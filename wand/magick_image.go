package wand

/*
#cgo LDFLAGS: -lMagickWand-Q16 -lMagickCore-Q16
#cgo CFLAGS: -fopenmp -I/usr/local/include/ImageMagick

#include <wand/magick_wand.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

/* Composite one image onto another at the specified offset. */
func (w *MagickWand) CompositeImage(srcWand *MagickWand, compose, x, y int) error {
	if C.MagickCompositeImage(w.wand, srcWand.wand, C.CompositeOperator(compose), C.ssize_t(x), C.ssize_t(y)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("CompositeImage() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Reads an image or image sequence from a blob. */
func (w *MagickWand) ReadImageBlob(blob []byte, length uint) error {
	if C.MagickReadImageBlob(w.wand, unsafe.Pointer(&blob[0]), C.size_t(length)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("ReadImageBlob() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Implements direct to memory image formats. */
func (w *MagickWand) GetImageBlob(length *uint) []byte {
	blobPtr := unsafe.Pointer(C.MagickGetImageBlob(w.wand, (*C.size_t)(unsafe.Pointer(length))))
	blob := C.GoBytes(blobPtr, C.int(int(*length)))
	return blob
}

/* Reads an image or image sequence. */
func (w *MagickWand) ReadImage(fileName string) error {
	if C.MagickReadImage(w.wand, C.CString(fileName)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("ReadImage() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Writes an image to the specified filename. */
func (w *MagickWand) WriteImage(fileName string) error {
	if C.MagickWriteImage(w.wand, C.CString(fileName)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("WriteImage() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Scales an image to the desired dimensions. */
func (w *MagickWand) ResizeImage(columns, rows uint, filter int, blur float64) error {
	if C.MagickResizeImage(w.wand, C.size_t(columns), C.size_t(rows), C.FilterTypes(filter), C.double(blur)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("Resize() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Sets the image compression quality. */
func (w *MagickWand) SetImageCompressionQuality(quality uint) error {
	if C.MagickSetImageCompressionQuality(w.wand, C.size_t(quality)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("SetImageCompressionQuality() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Gets the image compression quality. */
func (w *MagickWand) GetImageCompressionQuality() uint {
	return uint(C.MagickGetImageCompressionQuality(w.wand))
}

/* Sets the image compression. */
func (w *MagickWand) SetImageCompression(compression int) error {
	if C.MagickSetImageCompression(w.wand, C.CompressionType(compression)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("SetImageCompression() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Gets the image compression. */
func (w *MagickWand) GetImageCompression() int {
	return int(C.MagickGetImageCompression(w.wand))
}

/* Sets the format of a particular image in a sequence */
func (w *MagickWand) SetImageFormat(format string) error {
	if C.MagickSetImageFormat(w.wand, C.CString(format)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("SetImageFormat() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Gets the format of a particular image in a sequence. */
func (w *MagickWand) GetImageFormat() string {
	return C.GoString(C.MagickGetImageFormat(w.wand))
}

/* Gets the image height. */
func (w *MagickWand) GetImageHeight() uint {
	return uint(C.MagickGetImageHeight(w.wand))
}

/* Gets the image width. */
func (w *MagickWand) GetImageWidth() uint {
	return uint(C.MagickGetImageWidth(w.wand))
}

/* Extracts a region of the image. */
func (w *MagickWand) CropImage(width, height uint, x, y int) error {
	if C.MagickCropImage(w.wand, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("CropImage() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

/* Adaptively resize image with data dependent triangulation. */
func (w *MagickWand) AdaptiveResizeImage(columns, rows uint) error {
	if C.MagickAdaptiveResizeImage(w.wand, C.size_t(columns), C.size_t(rows)) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("AdaptiveResizeImage() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

// Sets the image background color.
func (w *MagickWand) SetImageBackgroundColor(bg *PixelWand) error {
	if C.MagickSetImageBackgroundColor(w.wand, bg.wand) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("SetImageBackgroundColor() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

// Returns the image background color.
func (w *MagickWand) GetImageBackgroundColor(bg *PixelWand) error {
	if C.MagickGetImageBackgroundColor(w.wand, bg.wand) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("GetImageBackgroundColor() failed : [%d] %s", eStr, eCode)
	}

	return nil
}

// Adds a blank image canvas of the specified size and background color to the 
// wand.
func (w *MagickWand) NewImage(cols, rows uint, bg *PixelWand) error {
	if C.MagickNewImage(w.wand, C.size_t(cols), C.size_t(rows), bg.wand) == C.MagickFalse {
		eStr, eCode := w.GetException()
		return fmt.Errorf("NewImage() failed : [%d] %s", eStr, eCode)
	}

	return nil
}
