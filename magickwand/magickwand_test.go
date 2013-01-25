package magickwand

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	Genesis()
	wand := New()
	wand.Destroy()
	Terminus()
}

func TestClear(t *testing.T) {
	Genesis()
	wand := New()
	wand.Clear()
	wand.Destroy()
	Terminus()
}

func TestReadBlob(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	file, err := os.Open("../examples/input/test.png")
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}

	defer file.Close()

	buf := &bytes.Buffer{}
	num, err := io.Copy(buf, file)
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}

	if err = wand.ReadBlob(buf.Bytes(), uint(num)); err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}
}

func TestWrite(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.Write("../examples/output/test.png"); err != nil {
		t.Error(err)
	}
}

func TestBlob(t *testing.T) {
	var length uint
	Genesis()
	defer Terminus()
	wand := New()
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

	if err = wand.ReadBlob(buf.Bytes(), uint(num)); err != nil {
		t.Errorf("Error: %s\n", err)
	}
	blob := wand.Blob(&length)

	file1, err := os.OpenFile("../examples/output/test.png",
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Error(err)
	}
	defer file1.Close()

	file1.Write(blob)
}

func TestCrop(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.Crop(100, 100, 100, 0); err != nil {
		t.Error(err)
	}

	if err := wand.Write("../examples/output/test-crop.png"); err != nil {
		t.Error(err)
	}
}

func TestAdaptiveResize(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.AdaptiveResize(800, 600); err != nil {
		t.Error(err)
	}

	if err := wand.Write("../examples/output/test-adaptiveresize.png"); err !=
		nil {
		t.Error(err)
	}
}

func TestResize(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.Resize(192, 108, GaussianFilter, 1.0); err != nil {
		t.Error(err)
	}

	if err := wand.Write("../examples/output/test-resize.png"); err !=
		nil {
		t.Error(err)
	}
}

func TestHeight(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.Height() != 1080 {
		t.Errorf("Height(%d) not equals 1080", wand.Height())
	}
}

func TestWidth(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.Width() != 1920 {
		t.Errorf("Width(%d) not equals 1080", wand.Width())
	}
}

func TestSetQuality(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.SetQuality(1); err != nil {
		t.Error(err)
	}

	if err := wand.Write("../examples/output/test-quality.png"); err !=
		nil {
		t.Error(err)
	}
}

func TestQuality(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.Quality() != 0 {
		t.Errorf("Quality(%d) not equals 100", wand.Quality())
	}

	if err := wand.Write("../examples/output/test-quality.png"); err !=
		nil {
		t.Error(err)
	}
}

func TestSetCompression(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.SetCompression(JPEGCompression); err != nil {
		t.Error(err)
	}

	if err := wand.Write("../examples/output/test-compression.jpg"); err !=
		nil {
		t.Error(err)
	}
}

func TestCompression(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.Compression() != ZipCompression {
		t.Error("Compression(%d) not equanls ZipCompression",
			int(wand.Compression()))
	}
}

func TestSetFormat(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if err := wand.SetFormat("JPEG"); err != nil {
		t.Error(err)
	}

	if err := wand.Write("../examples/output/test-format.jpg"); err !=
		nil {
		t.Error(err)
	}
}

func TestFormat(t *testing.T) {
	Genesis()
	defer Terminus()
	wand := New()
	defer wand.Destroy()

	if err := wand.Read("../examples/input/test.png"); err != nil {
		t.Error(err)
	}

	if wand.Format() != "PNG" {
		t.Errorf("Format(%s) not equanls PNG", wand.Format())
	}
}
