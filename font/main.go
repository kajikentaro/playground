package main

import (
	"image/color"
	"log"

	"github.com/fogleman/gg"
)

func main() {
	drawOtf()
	drawTtf()
}

func drawTtf() {
	dc := gg.NewContext(1200, 628)

	face := fetchTrueTypeFontFace("resource/GoNotoCJKCore.ttf")
	// PostScript は TrueTypeに対する後方互換性がある(?). ↓でも動く.
	// face := fetchPostScriptFontFace("resource/GoNotoCJKCore.ttf")
	dc.SetFontFace(face)
	dc.SetColor(color.White)
	dc.DrawStringWrapped("あいうえお", 50, 200, 0, 0, 600, 1, gg.AlignRight)

	if err := dc.SavePNG("test-ttf.png"); err != nil {
		log.Fatalln(err)
	}
}

func drawOtf() {
	dc := gg.NewContext(1200, 628)

	face := fetchPostScriptFontFace("resource/NotoSansJP-Medium.otf")
	dc.SetFontFace(face)
	dc.SetColor(color.White)
	dc.DrawStringWrapped("あいうえお", 50, 200, 0, 0, 600, 1, gg.AlignRight)

	if err := dc.SavePNG("test-otf.png"); err != nil {
		log.Fatalln(err)
	}
}
