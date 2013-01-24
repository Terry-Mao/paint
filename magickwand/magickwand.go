package magickwand

/*
#cgo LDFLAGS: -lMagickWand -lMagickCore
#cgo CFLAGS: -fopenmp -I/usr/include/ImageMagick

#include <wand/magick_wand.h>
*/
import "C"

import (
	"unsafe"
)

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
func New() *MagickWand {
    return &MagickWand { wand : C.NewMagickWand() }
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

func (w *MagickWand) Exception() string {
	var severity C.ExceptionType
	errPtr := C.MagickGetException(w.wand, &severity)
	C.MagickClearException(w.wand)
	err := C.GoString(errPtr)
	C.MagickRelinquishMemory(unsafe.Pointer(errPtr))

	return err
}
