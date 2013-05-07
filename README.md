## Terry-Mao/paint

`Terry-Mao/paint` is an image processing library based on ImageMagick for golang.

## Requeriments

ImageMagick's MagickWand development files are required.

```sh
# OSX
$ brew install imagemagick

# Arch Linux
$ sudo pacman -S extra/imagemagick

# Debian
$ sudo aptitude install libmagickwand-dev
```

## Installation

Just pull `Terry-Mao/paint` from github using `go get`:

```sh
$ go get github.com/Terry-Mao/paint
```

## Usage

```go
package main

import (
    "github.com/Terry-Mao/paint"
    "github.com/Terry-Mao/paint/wand"
)

func main() {
    wand.Genesis()
    defer wand.Terminus()
    w := wand.NewMagickWand()
    defer w.Destroy()

    if err := w.ReadImage("./examples/input/test2.jpg"); err != nil {
        t.Error(err)
    }

    if err := paint.Thumbnail(w, 302, 126); err != nil {
        t.Error(err)
    }

    if err := w.WriteImage("./examples/output/test2-thumbnail.jpg"); err != nil {
        t.Error(err)
    }

}
```

## Documentation

Read the `Terry-Mao/paint` documentation from a terminal

```go
$ go doc github.com/Terry-Mao/paint
$ go doc github.com/Terry-Mao/paint/magickwand
```

Alternatively, you can [paint](http://go.pkgdoc.org/github.com/Terry-Mao/paint) and [paint/wand](http://go.pkgdoc.org/github.com/Terry-Mao/paint/wand) online.
