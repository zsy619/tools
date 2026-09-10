package ximage

import (
	"image"
	"image/color"
	"image/draw"
)

// CircleMask 实现一个圆形的 image.Image 蒙版。
// 圆心为 P，半径为 R；圆内像素为不透明 Alpha，圆外为完全透明。
type CircleMask struct {
	P image.Point // 圆心坐标
	R int         // 半径
}

// ColorModel 返回蒙版的色彩模型，此处为 Alpha。
func (c *CircleMask) ColorModel() color.Model {
	return color.AlphaModel
}

// Bounds 返回蒙版覆盖的矩形区域。
func (c *CircleMask) Bounds() image.Rectangle {
	return image.Rect(c.P.X-c.R, c.P.Y-c.R, c.P.X+c.R, c.P.Y+c.R)
}

// At 返回坐标 (x, y) 处的颜色：圆内不透明，圆外透明。
func (c *CircleMask) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.P.X)+0.5, float64(y-c.P.Y)+0.5, float64(c.R)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// MergeImageNew 将 base 与 mask 叠加到一张大小与 base 相同的 RGBA 图像上。
// paddingX/paddingY 控制 mask 相对底板右下角的内边距。
// 当前实现固定使用 base 的尺寸作为输出尺寸。
func MergeImageNew(base image.Image, mask image.Image, paddingX int, paddingY int) (*image.RGBA, error) {
	baseSrcBounds := base.Bounds().Max
	maskSrcBounds := mask.Bounds().Max
	newWidth := baseSrcBounds.X
	newHeight := baseSrcBounds.Y
	maskWidth := maskSrcBounds.X
	maskHeight := maskSrcBounds.Y
	des := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight)) // 底板
	// 首先将底图绘制到目标画布
	draw.Draw(des, des.Bounds(), base, base.Bounds().Min, draw.Over)
	// 再将 mask 按偏移绘制到画布（右下角对齐）
	draw.Draw(des, image.Rect(paddingX, newHeight-paddingY-maskHeight, (paddingX+maskWidth), (newHeight-paddingY)), mask, image.Point{}, draw.Over)
	return des, nil
}
