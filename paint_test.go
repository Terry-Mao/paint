package paint

import (
	"github.com/Terry-Mao/paint/wand"
	"testing"
)

func TestThumbnail(t *testing.T) {

	wand.Genesis()
	defer wand.Terminus()
	wand := wand.NewMagickWand()
	defer wand.Destroy()

	if err := wand.ReadImage("./examples/input/test2.jpg"); err != nil {
		t.Error(err)
	}

	if err := Thumbnail(wand, 302, 126); err != nil {
		t.Error(err)
	}

	if err := wand.WriteImage("./examples/output/test2-thumbnail.jpg"); err != nil {
		t.Error(err)
	}
}

func BenchmarkThumbnail(b *testing.B) {
	wand.Genesis()
	defer wand.Terminus()
	wand := wand.NewMagickWand()
	defer wand.Destroy()

	b.StopTimer()
	b.StartTimer()

	for i := 0; i < 1000; i++ {
		if err := wand.ReadImage("./examples/input/test2.jpg"); err != nil {
			panic(err)
		}

		if err := Thumbnail(wand, 302, 126); err != nil {
			panic(err)
		}

		if err := wand.WriteImage("./examples/output/test2-thumbnail.jpg"); err != nil {
			panic(err)
		}

		wand.Clear()
	}
}
