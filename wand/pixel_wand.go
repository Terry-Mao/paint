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

// PixelWand wrap the struct of Imagemagick's PixelWand
type PixelWand struct {
	wand *C.PixelWand
}

// Returns a new pixel wand.
func NewPixelWand() *PixelWand {
	return &PixelWand{wand: C.NewPixelWand()}
}

// Clears resources associated with the wand.
func (p *PixelWand) Clear() {
	C.ClearPixelWand(p.wand)
}

// Deallocates resources associated with a PixelWand.
func (p *PixelWand) Destroy() {
	C.DestroyPixelWand(p.wand)
}

// Returns the severity, reason, and description of any error that occurs 
// when using other methods in this API.
func (p *PixelWand) Exception() (string, int) {
	var severity C.ExceptionType
	errPtr := C.PixelGetException(p.wand, &severity)
	C.PixelClearException(p.wand)
	err := C.GoString(errPtr)
	C.MagickRelinquishMemory(unsafe.Pointer(errPtr))
	return err, int(severity)
}

// Sets the color of the pixel wand with a string (e.g. "blue", "#0000ff", 
// "rgb(0,0,255)", "cmyk(100,100,100,10)", etc.).
func (p *PixelWand) SetColor(color string) error {
	if C.PixelSetColor(p.wand, C.CString(color)) == C.MagickFalse {
		eStr, eCode := p.Exception()
		return fmt.Errorf("SetColor() failed : [%d] %s", eStr, eCode)
    }

    return nil
}
