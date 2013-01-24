package magickimage

import (
	"testing"
    "os"
    "bytes"
    "io"
	"github.com/Terry-Mao/paint/magickwand"
)

func TestOpenBlob(t *testing.T) {
	magickwand.Genesis()
	wand := magickwand.New()
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

	err = wand.ReadBlob(buf.Bytes(), uint(num))
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}
	wand.Destroy()
	wand.Terminus()
}
