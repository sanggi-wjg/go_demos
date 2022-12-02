package barcode

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d/draw2dimg"
)

var RGBA_WHITE = color.RGBA{255, 255, 255, 255}

func CreateBarcode(code string) error {
	bcode, err := code128.Encode(code)
	if err != nil {
		return err
	}
	barcoded, _ := barcode.Scale(bcode, 250, 50)

	img := image.NewRGBA(image.Rect(0, 0, 250, 50))
	draw.Draw(img, img.Bounds(), &image.Uniform{RGBA_WHITE}, image.Pt(0, 0), draw.Src)
	gc := draw2dimg.NewGraphicContext(img)
	gc.FillStroke()
	gc.SetFillColor(image.Black)
	gc.SetFontSize(14)
	gc.FillStringAt(code, 50, 80)

	bcodeFile := imaging.New(270, 90, RGBA_WHITE)
	bcodeFile = imaging.Paste(bcodeFile, barcoded, image.Pt(10, 20))

	err = draw2dimg.SaveToPngFile(fmt.Sprintf("%s.png", code), bcodeFile)
	if err != nil {
		return err
	}

	return nil
}
