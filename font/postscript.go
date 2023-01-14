package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// PostScript アウトラインのフォント読み込み
func fetchPostScriptFontFace(fontfile string) font.Face {
	ftBinary, err := os.ReadFile(fontfile)
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

	return face
}
