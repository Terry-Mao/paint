package wand

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewMagickWand(t *testing.T) {
	Genesis()
	wand := NewMagickWand()
	wand.Destroy()
	Terminus()
}

func TestMagickWandClear(t *testing.T) {
	Genesis()
	wand := NewMagickWand()
	wand.Clear()
	wand.Destroy()
	Terminus()
}

func TestReadImageBlob(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	file, err := os.Open("../examples/input/test.png")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	buf := &bytes.Buffer{}
	num, err := io.Copy(buf, file)
	if err != nil {
		t.Error(err)
	}

	if err = wand.ReadImageBlob(buf.Bytes(), uint(num)); err != nil {
		t.Error(err)
	}
}

func TestReadImage(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}
}

func TestWriteImage(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("../examples/output/test.png"); err != nil {
		t.Error(err)
	}
}

func TestImageBlob(t *testing.T) {
	var length uint
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	file, err := os.Open("../examples/input/test.png")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	buf := &bytes.Buffer{}
	num, err := io.Copy(buf, file)
	if err != nil {
		t.Error(err)
	}

	if err = wand.ReadImageBlob(buf.Bytes(), uint(num)); err != nil {
		t.Error(err)
	}

	blob := wand.ImageBlob(&length)
	file1, err := os.OpenFile("../examples/output/test-blob.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Error(err)
	}
	defer file1.Close()

	file1.Write(blob)
}

func TestCropImage(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.CropImage(100, 100, 100, 0); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("../examples/output/test-crop.png"); err != nil {
		t.Error(err)
	}
}

func TestAdaptiveResizeImage(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.AdaptiveResizeImage(800, 600); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("../examples/output/test-adaptiveresize.png"); err != nil {
		t.Error(err)
	}
}

func TestResizeImage(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.ResizeImage(192, 108, GaussianFilter, 1.0); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("../examples/output/test-resize.png"); err != nil {
		t.Error(err)
	}
}

func TestImageHeight(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.ImageHeight() != 1080 {
		t.Errorf("Height(%d) not equals 1080", wand.ImageHeight())
	}
}

func TestImageWidth(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.ImageWidth() != 1920 {
		t.Errorf("Width(%d) not equals 1080", wand.ImageWidth())
	}
}

func TestSetImageCompressionQuality(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.SetImageCompressionQuality(1); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("../examples/output/test-quality.png"); err !=
		nil {
		t.Error(err)
	}
}

func TestImageCompressionQuality(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.ImageCompressionQuality() != 0 {
		t.Errorf("Quality(%d) not equals 100", wand.ImageCompressionQuality())
	}

	if err := wand.WriteImage("../examples/output/test-quality.png"); err != nil {
		t.Error(err)
	}
}

func TestSetImageCompression(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.SetImageCompression(JPEGCompression); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("../examples/output/test-compression.jpg"); err != nil {
		t.Error(err)
	}
}

func TestImageCompression(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.ImageCompression() != ZipCompression {
		t.Error("Compression(%d) not equanls ZipCompression", int(wand.ImageCompression()))
	}
}

func TestSetImageFormat(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.SetImageFormat("JPEG"); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("../examples/output/test-format.jpg"); err != nil {
		t.Error(err)
	}
}

func TestImageFormat(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.ImageFormat() != "PNG" {
		t.Errorf("Format(%s) not equanls PNG", wand.ImageFormat())
	}
}

func TestNewImage(t *testing.T) {
	Genesis()
	defer Terminus()
    wand := NewMagickWand()
	defer wand.Destroy()
	bg := NewPixelWand()
    defer bg.Destroy()
    if err := bg.SetColor("red"); err != nil {
        t.Error(err)
    }

    if err := wand.NewImage(300, 300, bg); err != nil {
        t.Error(err)
    }

	if err := wand.WriteImage("../examples/output/test-new-image.jpg"); err != nil {
		t.Error(err)
	}
}

func TestCompositeImage(t *testing.T) {
	Genesis()
	defer Terminus()
    wand := NewMagickWand()
	defer wand.Destroy()
	bg := NewPixelWand()
    defer bg.Destroy()
    if err := bg.SetColor("red"); err != nil {
        t.Error(err)
    }

    if err := wand.NewImage(300, 300, bg); err != nil {
        t.Error(err)
    }

	wand1 := NewMagickWand()
	defer wand1.Destroy()
	if err := wand1.ReadImage("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

    if err := wand1.CompositeImage(wand, AtopCompositeOp, 0, 0); err != nil {
        t.Error(err)
    }

	if err := wand1.WriteImage("../examples/output/test-composite-image.jpg"); err != nil {
		t.Error(err)
	}
}
