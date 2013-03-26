package paint

import (
	"github.com/Terry-Mao/paint/wand"
	"testing"
)

func TestThumbnail(t *testing.T) {

	wand.Genesis()
	defer wand.Terminus()
	w := wand.NewMagickWand()
	defer w.Destroy()

	if err := w.ReadImage("./examples/input/test2.jpg"); err != nil {
		t.Error(err)
	}

	if err := Thumbnail(w, 302, 126); err != nil {
		t.Error(err)
	}

	if err := w.WriteImage("./examples/output/test2-thumbnail.jpg"); err != nil {
		t.Error(err)
	}
}

func BenchmarkThumbnail(b *testing.B) {
	wand.Genesis()
	defer wand.Terminus()
	w := wand.NewMagickWand()
	defer w.Destroy()

	b.StopTimer()
	b.StartTimer()

	for i := 0; i < 1000; i++ {
		if err := w.ReadImage("./examples/input/test2.jpg"); err != nil {
			panic(err)
		}

		if err := Thumbnail(w, 302, 126); err != nil {
			panic(err)
		}

		if err := w.WriteImage("./examples/output/test2-thumbnail.jpg"); err != nil {
			panic(err)
		}

		w.Clear()
	}
}
