package main

import (
	"fmt"
	"image"
	"os"

	"github.com/signintech/pdft"
)

func main() {

	source := "source.pdf"
	signImage := "sign.png"
	dest := "dest.pdf"

	var pdf pdft.PDFt

	err := pdf.Open(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pages := pdf.GetNumberOfPage()

	pic, err := os.ReadFile(signImage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	infile, err := os.Open(signImage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ratio := float64(img.Bounds().Max.X) / float64(img.Bounds().Max.Y)
	w := 120.0
	h := w / ratio
	x := 160.0 //210.0 - float64(w)/2.0 + 300.0
	y := 480.0 //297.0 - float64(h)/2.0 + 50.0

	err = pdf.InsertImg(pic, pages, x, y, w, h)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pdf.Save(dest)
}
