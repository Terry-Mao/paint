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

import "github.com/Terry-Mao/paint"

func main() {
    magickwand.Genesis()
        defer magickwand.Terminus()
        wand := magickwand.New()
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
```

## Documentation

Read the `Terry-Mao/paint` documentation from a terminal

```go
$ go doc github.com/Terry-Mao/paint
$ go doc github.com/Terry-Mao/paint/magickwand
```

Alternatively, you can [paint](http://go.pkgdoc.org/github.com/Terry-Mao/paint) and [paint/magickwand](http://go.pkgdoc.org/github.com/Terry-Mao/paint/magickwand) online.
