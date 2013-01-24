package magickimage

/*
#cgo LDFLAGS: -lMagickWand -lMagickCore
#cgo CFLAGS: -fopenmp -I/usr/include/ImageMagick

#include <wand/magick_wand.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
	"github.com/Terry-Mao/paint/magickwand"
)

type MagickWand magickwand.MagickWand

func (w *MagickWand) magickwandReadBlob(blob []byte, length uint) error {
	if C.MagickReadImageBlob(w.wand, unsafe.Pointer(&blob[0]),
		C.size_t(length)) == C.MagickFalse {
		return fmt.Errorf("ReadBlob() failed : %s", magickwand.Exception(wand))
	}

	return nil
}
