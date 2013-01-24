package magickwand

import (
    "os"
    "io"
    "bytes"
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

func TestOpenBlob(t *testing.T) {
	Genesis()
	wand := New()
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
    Terminus()
}
