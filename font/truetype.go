package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func fetchTrueTypeFontFace(fontpath string) font.Face {
	ftBinary, err := os.ReadFile(fontpath)
	if err != nil {
		log.Fatalln(err)
	}

	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	opt := truetype.Options{
		Size:              90,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	face := truetype.NewFace(ft, &opt)
	return face
}
