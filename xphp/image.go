package xphp

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime"
	"os"
	"strings"

	gntext "github.com/g3n/engine/text"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/bmp"
	"golang.org/x/image/font"

	"haedu.gov.cn/tools/xstring"
)

// ImageInfo stores the info of an image
type ImageInfo struct {
	Width  int
	Height int
	Format string
	Mime   string
}

// GetImageSize gets the size of an image
func GetImageSize(filename string) (ImageInfo, error) {
	var info ImageInfo

	file, err := os.Open(filename)
	if err != nil {
		return info, err
	}
	defer file.Close()

	cfg, format, err := image.DecodeConfig(file)
	if err != nil {
		return info, err
	}

	info.Width = cfg.Width
	info.Height = cfg.Height
	info.Format = format
	info.Mime = mime.TypeByExtension("." + format)

	return info, nil
}

// GetImageSizeFromString gets the size of an image from a string
func GetImageSizeFromString(data []byte) (ImageInfo, error) {
	var info ImageInfo

	cfg, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return info, err
	}

	info.Width = cfg.Width
	info.Height = cfg.Height
	info.Format = format
	info.Mime = mime.TypeByExtension("." + format)

	return info, nil
}

var FontFamilies map[string]*truetype.Font

func init() {
	FontFamilies = make(map[string]*truetype.Font)
}

// 得到输入文本的区域大小
// 这几个变量分别是 字体大小, 角度, 字体名称, 字符串
// MeasureText returns the minimum width and height in pixels necessary for an image to contain
// the specified text. The supplied text string can contain line break escape sequences (\n).
func ImageTtfbBox(size, angle float64, fontPath string, text string) (with int, heght int) {
	fontx, err := gntext.NewFont(fontPath)
	if err != nil {
		return 0, 0
	}
	fontx.SetDPI(72)
	fontx.SetPointSize(size)
	fontx.SetHinting(font.HintingNone)
	return fontx.MeasureText(text)
	// fontFamily, ok := FontFamilies[fontPath]
	// if !ok {
	// 	fontFamily, _ := GetFontFamily(fontPath)
	// 	FontFamilies[fontPath] = fontFamily
	// }
	// f := freetype.NewContext()
	// // 设置用于绘制文本的字体
	// f.SetFont(fontFamily)
	// // 设置屏幕每英寸的分辨率
	// f.SetDPI(72)
	// f.SetHinting(font.HintingNone)
	// f.SetFontSize(size)
	// fixed := f.PointToFixed(size)
	// return float64(fixed)
}

// 获取字符集，仅调用一次
func GetFontFamily(fontPath string) (*truetype.Font, error) {
	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		fmt.Println("GetFontFamily--》", "read file error:", err)
		return &truetype.Font{}, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("GetFontFamily--》", "parse font error:", err)
		return &truetype.Font{}, err
	}
	return f, err
}

// gd库计算文本换行
func AutoWrap(size, angle float64, fontPath string, text string, width float64, line int) ([]string, int) {
	// 这几个变量分别是 字体大小, 角度, 字体名称, 字符串, 预设宽度, 行数
	content := ""
	height := 0
	// 将字符串拆分成一个个单字 保存到数组 letter 中
	letter := xstring.ToLetterArray(text)
	testLine := 1
	for _, l := range letter {
		teststr := content + " " + l
		testbox, xheight := ImageTtfbBox(size, 0, fontPath, teststr)
		height = xheight
		fmt.Println("testboxtestboxtestboxtestbox", testbox, width)
		// 判断拼接后的字符串是否超过预设的宽度
		if (float64(testbox) > width) && (content != "") {
			testLine++
			if testLine > line {
				runeArray := []rune(content)
				length := len(runeArray) - 2
				content = string(runeArray[:length])
				content += "......"
				break
			} else {
				content += PHP_EOL
			}
		}
		if testLine > line {
			break
		}
		content += l
	}
	return strings.Split(content, PHP_EOL), height
}

func LoadFontFace(path string, points float64) (font.Face, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size:    points, //
		Hinting: font.HintingFull,
	})
	return face, nil
}

// ImageJPEG jpeg图片输出到文件
func ImageJPEG(fpath string, m image.Image, quantity int) (file string, err error) {
	f, err := os.Create(fpath) // 创建文件
	if err != nil {
		fmt.Println("ImageJPEG:", err.Error())
		return "", err
	}
	defer f.Close() // 关闭文件
	if err != nil {
		fmt.Println("ImageJPEG:", err.Error())
		return "", err
	}
	if err = jpeg.Encode(f, m, &jpeg.Options{Quality: quantity}); err != nil { // 写入文件;
		fmt.Println("ImageJPEG:", err.Error())
		return "", err
	}

	file = fpath
	return
}

// ImagePNG png图片输出到文件
func ImagePNG(fpath string, m image.Image) (file string, err error) {
	f, err := os.Create(fpath) // 创建文件
	if err != nil {
		fmt.Println("ImagePNG:", err.Error())
		return "", err
	}
	defer f.Close() // 关闭文件
	if err != nil {
		fmt.Println("ImagePNG:", err.Error())
		return "", err
	}
	if err = png.Encode(f, m); err != nil { // 写入文件;
		fmt.Println("ImagePNG:", err.Error())
		return "", err
	}

	file = fpath
	return
}

// ImageGIF gif图片输出到文件
func ImageGIF(fpath string, m image.Image) (file string, err error) {
	f, err := os.Create(fpath) // 创建文件
	if err != nil {
		fmt.Println("ImageGIF:", err.Error())
		return "", err
	}
	defer f.Close() // 关闭文件
	if err != nil {
		fmt.Println("ImageGIF:", err.Error())
		return "", err
	}
	if err = gif.Encode(f, m, &gif.Options{}); err != nil { // 写入文件;
		fmt.Println("ImageGIF:", err.Error())
		return "", err
	}

	file = fpath
	return
}

// ImageBMP bmp图片输出到文件
func ImageBMP(fpath string, m image.Image) (file string, err error) {
	f, err := os.Create(fpath) // 创建文件
	if err != nil {
		fmt.Println("ImageBMP:", err.Error())
		return "", err
	}
	defer f.Close() // 关闭文件
	if err != nil {
		fmt.Println("ImageBMP:", err.Error())
		return "", err
	}
	if err = bmp.Encode(f, m); err != nil { // 写入文件;
		fmt.Println("ImageBMP:", err.Error())
		return "", err
	}

	file = fpath
	return
}

func DrawText(fontFamily *truetype.Font, text string, x, y int, fontSize float64, fontColor *image.Uniform, target *image.NRGBA) (*image.NRGBA, error) {
	f := freetype.NewContext()
	// 设置用于绘制文本的字体
	f.SetFont(fontFamily)
	// 设置屏幕每英寸的分辨率
	f.SetDPI(72)
	// 设置剪裁矩形以进行绘制
	f.SetClip(target.Bounds())
	// 设置目标图像
	f.SetDst(target)
	f.SetSrc(fontColor)
	// 以磅为单位设置字体大小
	f.SetFontSize(fontSize)
	f.SetHinting(font.HintingNone)
	// 获取字体的尺寸大小
	// fixed := f.PointToFixed(fontSize)
	// pt := freetype.Pt(x-(utf8.RuneCountInString(text)/2)*fixed.Ceil(), y)
	pt := freetype.Pt(x, y)
	// 根据 Pt 的坐标值绘制给定的文本内容
	_, err := f.DrawString(text, pt)
	if err != nil {
		fmt.Println("DrawText:", err.Error())
		return nil, err
	}
	return target, nil
}

func DrawTextMultiline(fontFamily *truetype.Font, text []string, height int, x, y int, fontSize float64, maxWidth float64, fontColor *image.Uniform, target *image.NRGBA) (*image.NRGBA, error) {
	f := freetype.NewContext()
	// 设置用于绘制文本的字体
	f.SetFont(fontFamily)
	// 设置屏幕每英寸的分辨率
	f.SetDPI(72)
	// 设置剪裁矩形以进行绘制
	f.SetClip(target.Bounds())
	// 设置目标图像
	f.SetDst(target)
	f.SetSrc(fontColor)
	// 以磅为单位设置字体大小
	f.SetFontSize(fontSize)
	f.SetHinting(font.HintingNone)
	// 获取字体的尺寸大小
	for i, s := range text {
		posY := y
		if i != 0 {
			posY = y + i*height
		}
		pt := freetype.Pt(x, posY)
		// 根据 Pt 的坐标值绘制给定的文本内容
		_, err := f.DrawString(s, pt)
		if err != nil {
			fmt.Println("DrawTextMultiline:", err.Error())
			return nil, err
		}
	}
	return target, nil
}
