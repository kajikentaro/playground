package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func main() {
	flag.Parse()

	dc := gg.NewContext(1200, 628)

	backgroundImage, err := gg.LoadImage("src/kajindowsIcon.png")
	if err != nil {
		panic(errors.Wrap(err, "load background image"))
	}
	backgroundImage = imaging.Fill(backgroundImage, dc.Width(), dc.Height(), imaging.Center, imaging.Lanczos)
	dc.DrawImage(backgroundImage, 0, 0)

	margin := 12.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 127})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()

	// START
	ftBinary, err := os.ReadFile("./font/NotoSansJP-Medium.otf")
	if err != nil {
		log.Fatalln(err)
	}

	ft, err := opentype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	opt := opentype.FaceOptions{
		Size:    160.0,
		DPI:     72.0,
		Hinting: font.HintingNone,
	}

	face, _ := opentype.NewFace(ft, &opt)
	// END

	dc.SetFontFace(face)
	dc.SetColor(color.White)
	s := "あいうえお"
	maxWidth := float64(dc.Width())
	dc.DrawStringWrapped(s, 50, 200, 0, 0, maxWidth/2, 1.5, gg.AlignLeft)

	if err := dc.SavePNG("test.png"); err != nil {
		panic(errors.Wrap(err, "save png"))
	}
}
