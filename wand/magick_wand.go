// Package wand providers Imagemagick functions.
package wand

/*
#cgo LDFLAGS: -lMagickWand-Q16 -lMagickCore-Q16
#cgo CFLAGS: -fopenmp -I/usr/local/include/ImageMagick

#include <wand/magick_wand.h>
*/
import "C"

import (
	"unsafe"
)

/* MagickWand wrap the struct of Imagemagick's MagickWand */
type MagickWand struct {
	wand *C.MagickWand
}

/* Initializes the MagickWand environment. */
func Genesis() {
	C.MagickWandGenesis()
}

/* Terminates the MagickWand environment. */
func Terminus() {
	C.MagickWandTerminus()
}

/* 
   Create a wand required for all other methods in the API. A panic() is 
   thrown if there is not enough memory to allocate the wand. Use Destroy() to 
   dispose of the wand when it is no longer needed. 
*/
func NewMagickWand() *MagickWand {
	return &MagickWand{wand: C.NewMagickWand()}
}

/* 
   clears resources associated with the wand, leaving the wand blank, and 
   ready to be used for a new set of images. 
*/
func (w *MagickWand) Clear() {
	C.ClearMagickWand(w.wand)
}

/* Deallocates memory associated with an MagickWand. */
func (w *MagickWand) Destroy() {
	C.DestroyMagickWand(w.wand)
}

/* Returns the severity, reason, and description of any error that occurs when 
   using other methods in this API. 
*/
func (w *MagickWand) GetException() (string, int) {
	var severity C.ExceptionType
	errPtr := C.MagickGetException(w.wand, &severity)
	C.MagickClearException(w.wand)
	err := C.GoString(errPtr)
	C.MagickRelinquishMemory(unsafe.Pointer(errPtr))
	return err, int(severity)
}

// Resets the wand iterator
func (w *MagickWand) ResetIterator() {
    C.MagickResetIterator(w.wand)
}
