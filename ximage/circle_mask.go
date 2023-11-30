package ximage

import (
	"image"
	"image/color"
	"image/draw"
)

type CircleMask struct {
	P image.Point
	R int
}

func (c *CircleMask) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *CircleMask) Bounds() image.Rectangle {
	return image.Rect(c.P.X-c.R, c.P.Y-c.R, c.P.X+c.R, c.P.Y+c.R)
}

func (c *CircleMask) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.P.X)+0.5, float64(y-c.P.Y)+0.5, float64(c.R)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// 图片拼接
func MergeImageNew(base image.Image, mask image.Image, paddingX int, paddingY int) (*image.RGBA, error) {
	baseSrcBounds := base.Bounds().Max
	maskSrcBounds := mask.Bounds().Max
	newWidth := baseSrcBounds.X
	newHeight := baseSrcBounds.Y
	maskWidth := maskSrcBounds.X
	maskHeight := maskSrcBounds.Y
	des := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight)) // 底板
	// 首先将一个图片信息存入jpg
	draw.Draw(des, des.Bounds(), base, base.Bounds().Min, draw.Over)
	// 将另外一张图片信息存入jpg
	// draw.Draw(des, image.Rect(paddingX, newHeight-paddingY-maskHeight, (paddingX+maskWidth), (newHeight-paddingY)), mask, image.ZP, draw.Over)
	draw.Draw(des, image.Rect(paddingX, newHeight-paddingY-maskHeight, (paddingX+maskWidth), (newHeight-paddingY)), mask, image.Point{}, draw.Over)
	return des, nil
}
