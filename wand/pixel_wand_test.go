
package wand

import (
	"testing"
)

func TestNewPixelWand(t *testing.T) {
	Genesis()
	wand := NewPixelWand()
	wand.Destroy()
	Terminus()
}

func TestPixelClear(t *testing.T) {
	Genesis()
	wand := NewPixelWand()
	wand.Clear()
	wand.Destroy()
	Terminus()
}

func TestSetColor(t *testing.T) {
	Genesis()
    defer Terminus()
	wand := NewPixelWand()
	defer wand.Destroy()
    if err := wand.SetColor("white"); err != nil {
		t.Error(err)
    }

	wand.Clear()
}
