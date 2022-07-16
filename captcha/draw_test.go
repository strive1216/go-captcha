/**
 * @Author Awen
 * @Description
 * @Date 2021/7/20
 **/

package captcha

import (
	_ "embed"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func GetDraw() *Draw {
	return &Draw{}
}

//go:embed resources/fonts/851手書き雑フォント.ttf
var _font []byte

//go:embed resources/images/1.jpg
var _img []byte

func TestDrawTextImg(t *testing.T) {
	draw := GetDraw()

	drawDots := DrawDot{
		Dx:      0,
		Dy:      0,
		FontDPI: 72,
		Text:    "你好",
		Angle:   45,
		Size:    20,
		Color:   "#841524",
		Width:   20,
		Height:  20,
		Font:    _font,
	}

	canvas, ap, _ := draw.DrawTextImg(drawDots, DrawCanvas{
		Width:      20,
		Height:     20,
		Background: _img,
	})

	nW := canvas.Bounds().Max.X
	nH := canvas.Bounds().Max.Y
	minX := ap.MinX
	maxX := ap.MaxX
	minY := ap.MinY
	maxY := ap.MaxY
	width := maxX - minX
	height := maxY - minY

	co, _ := ParseHexColor("#841524")
	coArr := []color.RGBA{
		co,
	}
	canvas2 := draw.CreateCanvasWithPalette(DrawCanvas{
		Width:  width,
		Height: height,
	}, coArr)

	// 开始裁剪
	for x := 0; x < nW; x++ {
		for y := 0; y < nH; y++ {
			co := canvas.At(x, y)
			if _, _, _, a := co.RGBA(); a > 0 {
				canvas2.Set(x, y, canvas.At(x, y))
			}
		}
	}

	file := fmt.Sprintf("%v", RandInt(1, 200)) + "textImg.png"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	defer logFile.Close()
	err := png.Encode(logFile, canvas)
	if err != nil {
		panic(err)
	}
}
