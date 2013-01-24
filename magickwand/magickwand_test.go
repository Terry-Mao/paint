package magickwand

import (
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
