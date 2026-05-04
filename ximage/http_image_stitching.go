package ximage

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"math"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

type HttpImageStitching struct{}

// readImgData 从url中获取图片资源
func (his *HttpImageStitching) readImgData(url string) image.Image {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("图片获取失败", err)
		return nil
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		fmt.Println("图片decode失败", err)
		return nil
	}
	return img
}

// calculateRatioFit 计算图片缩放后的尺寸
func (his *HttpImageStitching) calculateRatioFit(srcWidth, srcHeight int, defaultWidth, defaultHeight float64) (int, int) {
	ratio := math.Min(defaultWidth/float64(srcWidth), defaultHeight/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// ImageStitching 3图合1，创建图片
func (his *HttpImageStitching) ImageStitching(img1Url, img2Url, img3Url string) (map[string]interface{}, bool) {
	returnData := make(map[string]interface{})
	if img1Url == "" || img2Url == "" || img3Url == "" {
		return nil, false
	}
	// 根据图片地址获取图片.
	img1 := his.readImgData(img1Url)
	img2 := his.readImgData(img2Url)
	img3 := his.readImgData(img3Url)
	if img1 == nil || img2 == nil || img3 == nil {
		return nil, false
	}
	// 图片1缩放
	b1 := img1.Bounds()
	img1Width := b1.Max.X
	img1Height := b1.Max.Y
	w1, h1 := his.calculateRatioFit(img1Width, img1Height, 344, 344)
	img1m := resize.Resize(uint(w1), uint(h1), img1, resize.Lanczos3)
	// 图片2缩放
	b2 := img2.Bounds()
	img2Width := b2.Max.X
	img2Height := b2.Max.Y
	w2, h2 := his.calculateRatioFit(img2Width, img2Height, 172, 172)
	img2m := resize.Resize(uint(w2), uint(h2), img2, resize.Lanczos3)
	// 图片3缩放
	b3 := img3.Bounds()
	img3Width := b3.Max.X
	img3Height := b3.Max.Y
	w3, h3 := his.calculateRatioFit(img3Width, img3Height, 172, 172)
	img3m := resize.Resize(uint(w3), uint(h3), img3, resize.Lanczos3)
	// 创建源图
	fileName := "dst.jpg"
	file, err := os.Create(fileName)
	if err != nil {
		return nil, false
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("CreateGoodsPicture:图片资源关闭错误", err)
		}
	}()
	// 图片三合一绘图
	jpg := image.NewRGBA(image.Rect(0, 0, 344, 516))
	draw.Draw(jpg, jpg.Bounds().Add(image.Pt(0, 0)), img1m, img1m.Bounds().Min, draw.Src)
	draw.Draw(jpg, jpg.Bounds().Add(image.Pt(0, 344)), img2m, img2m.Bounds().Min, draw.Src)
	draw.Draw(jpg, jpg.Bounds().Add(image.Pt(172, 344)), img3m, img3m.Bounds().Min, draw.Src)
	// jpeg.Encode默认图片质量75%
	err1 := jpeg.Encode(file, jpg, nil)
	if err1 != nil {
		fmt.Println("CreateGoodsPicture:图片png.Encode错误", err1)
		return nil, false
	}

	defer func() {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("CreateGoodsPicture:图片资源删除错误", err)
		}
	}()
	return returnData, true
}
