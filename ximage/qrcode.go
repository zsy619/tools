package ximage

import (
	"crypto/md5"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	qrcode "github.com/skip2/go-qrcode"
	"golang.org/x/image/bmp"
	"golang.org/x/image/font"

	"github.com/zsy619/tools/xcrypto"
	"github.com/zsy619/tools/xio"
	"github.com/zsy619/tools/xphp"
)

//go:generate goption -p . -c QrCode -w
//go:generate goption -p . -c BackgroundConfig -w
//go:generate goption -p . -c QrCodeConfig -w
//go:generate goption -p . -c TitleConfig -w
//go:generate gofmt -w .

var (
	// Black is an opaque black uniform image.
	Black = image.NewUniform(color.Black)
	// White is an opaque white uniform image.
	White = image.NewUniform(color.White)
	// Transparent is a fully transparent uniform image.
	Transparent = image.NewUniform(color.Transparent)
	// Opaque is a fully opaque uniform image.
	Opaque = image.NewUniform(color.Opaque)
)

type QrCode struct {
	IsOne        bool             // 生成一次，按月存储，文件名 md5(内容)
	Root         string           // 最终图片保存路径
	Title        TitleConfig      // 标题配置
	Background   BackgroundConfig // 背景图片设置
	QrCodeConfig QrCodeConfig     // 二维码配置
}

type BackgroundConfig struct {
	Path   string // 背景图片路径
	Resize Resize // 背景缩放尺寸
}

type QrCodeConfig struct {
	LogoPath   string               // 二维码logo路径，在二维码上居中显示
	LogoResize Resize               // 二维码logo缩放尺寸
	Content    string               // 二维码内容
	Level      qrcode.RecoveryLevel // 二维码等级
	Size       int                  // 二维码大小
	Offset     Offset               // 二维码相对背景图片偏移位置
}

type TitleConfig struct {
	Title   string  // 标题
	Offset  Offset  // 标题位置
	Height  int     // 区域高度
	Hinting string  // none | full
	Dpi     float64 // screen resolution in Dots Per Inch
	Size    float64 // font size in points
	Spacing float64 //line spacing (e.g. 2 means double spaced)
	Path    string  // filename of the ttf font
	Wonb    bool    // white text on a black background
}

type Resize struct {
	Width  int // 宽度
	Height int // 高度
}

type Offset struct {
	X int
	Y int
}

// NewQrCodeDefault
/**
 * @description: 创建默认的二维码配置
 * @return {*}
 */
func NewQrCodeDefault() *QrCode {
	qrCode := &QrCode{
		Title: TitleConfig{
			Dpi:     72,
			Size:    42,
			Hinting: "none",
			Spacing: 1.5,
			Wonb:    false,
			Height:  200,
		},
		Background: BackgroundConfig{
			Resize: Resize{
				Width:  0,
				Height: 0,
			},
		},
		QrCodeConfig: QrCodeConfig{
			Offset: Offset{
				X: 0,
				Y: 0,
			},
			LogoResize: Resize{
				Width:  0,
				Height: 0,
			},
			Size:  620,
			Level: qrcode.Highest,
		},
	}

	return qrCode
}

/** CreateQrCodeBackground
 * @description: 生成带带背景图片的二维码
 * @param {*}
 * @return {string} 生成的图片路径
 */
func (code *QrCode) CreateQrCodeBackground() (file string, err error) {
	err = code.checkAttribute() // 参数校验
	if err != nil {
		return "", err
	}
	var (
		qrcode      image.Image
		offset      image.Point
		original    *os.File
		originalImg image.Image
	)

	// 二维码
	qrcode, err = code.createQrCode(code.QrCodeConfig.Content)
	if err != nil {
		return
	}

	// 背景图片
	original, err = os.Open(code.Background.Path)
	if err != nil {
		fmt.Println("", err.Error())
		return
	}
	defer original.Close()

	originalImg, err = png.Decode(original)
	if err != nil {
		fmt.Println("", err.Error())
		return
	}
	// 判断是否缩放
	if code.Background.Resize.Width > 0 && code.Background.Resize.Height > 0 {
		originalImg = code.ImageResize(originalImg, code.Background.Resize)
	}

	qrcodeB := qrcode.Bounds()
	originalB := originalImg.Bounds()
	if qrcodeB.Dx() >= originalB.Dx() || qrcodeB.Dy() >= originalB.Dy() {
		err = errors.New("背景图尺寸过小，至少为（宽*高）：")
		return
	}
	// 设置为居中
	offset = image.Pt((originalB.Max.X-qrcodeB.Max.X)/2+code.QrCodeConfig.Offset.X, (originalB.Max.Y-qrcodeB.Max.Y)/2+code.QrCodeConfig.Offset.Y)
	m := image.NewNRGBA(originalB)
	// draw.Draw(m, originalB, originalImg, image.ZP, draw.Src)
	draw.Draw(m, originalB, originalImg, image.Point{}, draw.Src)
	// draw.Draw(m, qrcodeB.Add(offset), qrcode, image.ZP, draw.Src)
	draw.Draw(m, qrcodeB.Add(offset), qrcode, image.Point{}, draw.Src)

	// 字体设置
	if code.Title.Title != "" {
		dst, errx := code.showTitle(m)
		if errx == nil {
			m = dst
		}
	}

	fpath := code.Root + code.randomFileName() + ".png"
	file, err = xphp.ImagePNG(fpath, m)
	return
}

// CreateQrCodeBackgroundPlacenum
/**
 * @description:
 * @param {string} file
 * @return {string , error}
 */
func (code *QrCode) CreateQrCodeBackgroundPlacenum(file string) (dest string, err error) {
	err = code.checkAttribute() // 参数校验
	if err != nil {
		return "", err
	}
	var (
		qrcode      image.Image
		offset      image.Point
		original    *os.File
		originalImg image.Image
	)

	// 二维码
	qrcode, err = code.createQrCode(code.QrCodeConfig.Content)
	if err != nil {
		fmt.Println("CreateQrCodeBackgroundPlacenum：", err.Error())
		return
	}

	// 背景图片
	original, err = os.Open(code.Background.Path)
	if err != nil {
		fmt.Println("CreateQrCodeBackgroundPlacenum：", err.Error())
		return
	}
	defer original.Close()

	originalImg, err = png.Decode(original)
	if err != nil {
		fmt.Println("CreateQrCodeBackgroundPlacenum：", err.Error())
		return
	}
	// 判断是否缩放
	if code.Background.Resize.Width > 0 && code.Background.Resize.Height > 0 {
		originalImg = code.ImageResize(originalImg, code.Background.Resize)
	}

	qrcodeB := qrcode.Bounds()
	originalB := originalImg.Bounds()
	if qrcodeB.Dx() >= originalB.Dx() || qrcodeB.Dy() >= originalB.Dy() {
		err = fmt.Errorf("背景图尺寸过小，至少为（宽*高）：%d*%d", qrcodeB.Dx(), qrcodeB.Dy())
		return
	}
	// 设置为居中
	offset = image.Pt((originalB.Max.X-qrcodeB.Max.X)/2+code.QrCodeConfig.Offset.X, (originalB.Max.Y-qrcodeB.Max.Y)/2+code.QrCodeConfig.Offset.Y)
	m := image.NewNRGBA(originalB)
	// draw.Draw(m, originalB, originalImg, image.ZP, draw.Src)
	// draw.Draw(m, qrcodeB.Add(offset), qrcode, image.ZP, draw.Over)
	draw.Draw(m, originalB, originalImg, image.Point{}, draw.Src)
	draw.Draw(m, qrcodeB.Add(offset), qrcode, image.Point{}, draw.Over)

	// 字体设置
	if code.Title.Title != "" {
		dst, errx := code.showTitle(m)
		if errx == nil {
			m = dst
		}
	}

	dest, err = xphp.ImagePNG(file, m)
	return
}

// showTitle
/**
 * @description: 显示标题
 * @param {*image.NRGBA} target
 * @return {*}
 */
func (code *QrCode) showTitle(target *image.NRGBA) (*image.NRGBA, error) {
	fontFamily, err := xphp.GetFontFamily(code.Title.Path)
	if err != nil {
		fmt.Println("showTitle：", "get font family error")
		return nil, err
	}
	f := freetype.NewContext()
	// 设置用于绘制文本的字体
	f.SetFont(fontFamily)
	// 设置屏幕每英寸的分辨率
	f.SetDPI(code.Title.Dpi)
	// 设置剪裁矩形以进行绘制
	f.SetClip(target.Bounds())
	// 设置目标图像
	f.SetDst(target)
	switch code.Title.Hinting {
	default:
		f.SetHinting(font.HintingNone)
	case "full":
		f.SetHinting(font.HintingFull)
	}
	// 设置绘制操作的源图像，通常为 image.Uniform
	// f.SetSrc(image.NewUniform(color.RGBA{R: 220, G: 220, B: 220, A: 220}))
	f.SetSrc(Black)
	// 以磅为单位设置字体大小
	f.SetFontSize(code.Title.Size)

	drawStr := code.Title.Title
	if drawStr == "" {
		drawStr = "嘿，世界！"
	}
	// 获取字体的尺寸大小
	fixed := f.PointToFixed(code.Title.Size)
	// fixed.Ceil() 字体大小
	// utf8.RuneCountInString(drawStr) 获取字符串的实际大小，而不是以byte算
	pt := freetype.Pt(target.Rect.Max.X/2-(utf8.RuneCountInString(drawStr)/2)*fixed.Ceil()+code.Title.Offset.X, code.Title.Offset.Y)
	// pt := freetype.Pt(0, code.Title.Offset.Y)
	// 根据 Pt 的坐标值绘制给定的文本内容
	fix, err := f.DrawString(drawStr, pt)
	if err != nil {
		fmt.Println("showTitle：", err.Error())
		return nil, err
	}
	fmt.Println("showTitle：", fix)
	return target, nil
}

// getFontFamily
/**
 * @description: 获取字符集，仅调用一次
 * @return {*truetype.Font, error}
 */
func (code *QrCode) getFontFamily() (*truetype.Font, error) {
	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := os.ReadFile(code.Title.Path)
	if err != nil {
		fmt.Println("getFontFamily--》", "read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("getFontFamily--》", "parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}

// CreateQrCode
/**
 * @description: 生成二维码
 * @param {string} content 内容
 * @param {int} size 大小
 * @param {qrcode.RecoveryLevel} level 二维码质量
 * @return {string , error}
 */
func (code *QrCode) CreateQrCode(content string, size int, level qrcode.RecoveryLevel) (file string, err error) {
	if content == "" {
		return "", errors.New("无二维码内容")
	}
	if code.IsOne {
		code.Root = "./Downloads/temp/qrcode_" + time.Now().Format("20060102") + "/"
		if err := os.MkdirAll(code.Root, os.ModePerm); err != nil {
			fmt.Println("CreateQrCode：", err.Error())
			return "", err
		}
		fpath := code.Root + code.md5FileName() + ".png"
		if xio.FileIsExisted(fpath) {
			return fpath, nil
		}
	}
	if size <= 0 {
		size = 620
	}
	code.QrCodeConfig.Size = size
	code.QrCodeConfig.Level = level
	if xphp.Empty(code.Root) {
		if code.IsOne {
			code.Root = "./Downloads/temp/qrcode_" + time.Now().Format("20060102") + "/"
		} else {
			code.Root = "./Downloads/temp/qrcode_" + time.Now().Format("20060102150405") + "/"
		}
	}
	if err := os.MkdirAll(code.Root, os.ModePerm); err != nil {
		fmt.Println("CreateQrCode：", err.Error())
		return "", err
	}
	img, err := code.createQrCode(content)
	if err != nil {
		fmt.Println("CreateQrCode：", err.Error())
		return "", err
	}

	fpath := code.randomFileName()
	if code.IsOne {
		fpath = code.md5FileName()
	}
	fpath = code.Root + fpath + ".png"
	file, err = xphp.ImagePNG(fpath, img)
	return
}

func (code *QrCode) CreateQrCodeImage(content string, size int, level qrcode.RecoveryLevel) (image.Image, error) {
	code.QrCodeConfig.Size = size
	code.QrCodeConfig.Level = level
	return code.createQrCode(content)
}

func (code *QrCode) CreateQrCodeImageByHighest(content string, size int) (image.Image, error) {
	return code.CreateQrCodeImage(content, size, qrcode.Highest)
}

func (code *QrCode) CreateQrCodeFileName(content string, size int, level qrcode.RecoveryLevel, file string) (ferr error) {
	if content == "" {
		return errors.New("无二维码内容")
	}
	if size <= 0 {
		size = 620
	}
	img, err := code.createQrCode(content)
	if err != nil {
		fmt.Println("CreateQrCode：", err.Error())
		return err
	}
	code.QrCodeConfig.Size = size
	code.QrCodeConfig.Level = level
	_, ferr = xphp.ImagePNG(file, img)
	return
}

// checkAttribute 属性校验
func (code *QrCode) checkAttribute() error {
	if code.QrCodeConfig.Content == "" {
		return errors.New("无二维码内容")
	}
	if code.Background.Path == "" {
		return errors.New("未设置背景图")
	}
	if code.QrCodeConfig.Size <= 0 {
		code.QrCodeConfig.Size = 620
	}
	if xphp.Empty(code.Root) {
		code.Root = "./Downloads/temp/qrcode_" + time.Now().Format("20060102150405") + "/"
	}
	if err := os.MkdirAll(code.Root, os.ModePerm); err != nil {
		fmt.Println("CreateQrCode：", err.Error())
		return err
	}
	return nil
}

// randomFileName
/**
 * @description: 获取一个随机的文件名称
 * @return {string}
 */
func (code *QrCode) randomFileName() string {
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
	fileName := fmt.Sprintf("%x", hashName)
	return fileName
}

// md5FileName
/**
 * @description: 获取一个md5文件名称
 * @return {string}
 */
func (code *QrCode) md5FileName() string {
	fileName := xcrypto.GetMD5Hash(code.QrCodeConfig.Content)
	return fileName
}

// createQrCode
/**
 * @description: 创建二维码
 * @param {string} content
 * @return {image.Image, error}
 */
func (code *QrCode) createQrCode(content string) (qrcodeImg image.Image, err error) {
	var qrCode *qrcode.QRCode
	qrCode, err = qrcode.New(content, code.QrCodeConfig.Level)
	if err != nil {
		return nil, errors.New("创建二维码失败:" + err.Error())
	}
	qrCode.DisableBorder = true
	qrcodeImg = qrCode.Image(code.QrCodeConfig.Size)

	// 二维码logo
	if code.QrCodeConfig.LogoPath != "" {
		logo, errx := os.Open(code.QrCodeConfig.LogoPath)
		if errx != nil {
			fmt.Println("createQrCode-->", errx.Error())
			goto OK
		}
		defer logo.Close()
		var logoImg image.Image
		ext := code.getFileExt(code.QrCodeConfig.LogoPath)
		fmt.Println("ext:", ext)
		switch ext {
		default:
			logoImg, errx = jpeg.Decode(logo)
		case "bmp":
			logoImg, errx = bmp.Decode(logo)
		case "png":
			logoImg, errx = png.Decode(logo)
		case "jpg":
			fallthrough
		case "jpeg":
			logoImg, errx = jpeg.Decode(logo)
		case "gif":
			logoImg, errx = gif.Decode(logo)
		}
		if errx != nil {
			fmt.Println("createQrCode-->", errx.Error())
			goto OK
		}
		// 判断是否缩放
		if code.QrCodeConfig.LogoResize.Width > 0 && code.QrCodeConfig.LogoResize.Height > 0 {
			logoImg = code.ImageResize(logoImg, code.QrCodeConfig.LogoResize)
		}

		qrcodeB := qrcodeImg.Bounds()
		logoB := logoImg.Bounds()
		// 设置为居中
		offset := image.Pt((qrcodeB.Dx()-logoB.Dx())/2, (qrcodeB.Dy()-logoB.Dy())/2)
		m := image.NewNRGBA(qrcodeB)
		// draw.Draw(m, qrcodeB, qrcodeImg, image.ZP, draw.Src)
		// draw.Draw(m, logoB.Add(offset), logoImg, image.ZP, draw.Src)
		draw.Draw(m, qrcodeB, qrcodeImg, image.Point{}, draw.Src)
		draw.Draw(m, logoB.Add(offset), logoImg, image.Point{}, draw.Src)

		qrcodeImg = m
	}
OK:
	return qrcodeImg, nil
}

// ImageResize
/**
 * @description:
 * @param {image.Image} src
 * @param {Resize} rse
 * @return {image.Image}
 */
func (code *QrCode) ImageResize(src image.Image, rse Resize) image.Image {
	return resize.Resize(uint(rse.Width), uint(rse.Height), src, resize.Lanczos3)
}

// checkFile
/**
 * @description: 检查文件是否存在
 * @param {string} name
 * @return {bool, error}
 */
func (code *QrCode) checkFile(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// getFileExt
/**
 * @description: 获取文件后缀
 * @param {string} filename
 * @return {string}
 */
func (code *QrCode) getFileExt(filename string) string {
	return strings.ToLower(strings.ReplaceAll(filepath.Ext(filename), ".", ""))
}
