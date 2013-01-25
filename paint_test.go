package paint

import (
	"github.com/Terry-Mao/paint/magickwand"
	"testing"
)

func TestThumbnail(t *testing.T) {

	magickwand.Genesis()
	defer magickwand.Terminus()
	wand := magickwand.New()
	defer wand.Destroy()

	if err := wand.Read("./examples/input/test2.jpg"); err != nil {
		t.Error(err)
	}

	if err := Thumbnail(wand, 302, 126); err != nil {
		t.Error(err)
	}

	if err := wand.Write("./examples/output/test2-thumbnail.jpg"); err != nil {
		t.Error(err)
	}
}

func BenchmarkThumbnail(b *testing.B) {
	magickwand.Genesis()
	defer magickwand.Terminus()
	wand := magickwand.New()
	defer wand.Destroy()

	b.StopTimer()
	b.StartTimer()

	for i := 0; i < 1000; i++ {
		if err := wand.Read("./examples/input/test2.jpg"); err != nil {
			panic(err)
		}

		if err := Thumbnail(wand, 302, 126); err != nil {
			panic(err)
		}

		if err := wand.Write("./examples/output/test2-thumbnail.jpg"); err !=
			nil {
			panic(err)
		}

		wand.Clear()
	}

}
