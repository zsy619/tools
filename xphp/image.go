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

	"github.com/zsy619/tools/xstring"
)

// ImageInfo 描述一张图片的基础信息。
type ImageInfo struct {
	Width  int    // 像素宽度
	Height int    // 像素高度
	Format string // 解码格式（png/jpeg/gif/bmp 等）
	Mime   string // 对应的 MIME 类型
}

// GetImageSize 读取 filename 指向的图片文件，返回其尺寸与格式信息。
// 打开文件失败或图片格式无法识别时返回相应错误。
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

// GetImageSizeFromString 从内存中的图片字节流解析尺寸与格式。
// 数据非合法图片或解码失败时返回相应错误。
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

// FontFamilies 字体文件路径到解析后 *truetype.Font 的缓存。
var FontFamilies map[string]*truetype.Font

func init() {
	FontFamilies = make(map[string]*truetype.Font)
}

// ImageTtfbBox 测量渲染 text 所需的最小宽度与高度（像素）。
//
// 参数依次为：字体大小（磅）、旋转角度（当前未使用）、字体文件路径、文本内容。
// 文本中可包含 \n 换行符；字体加载失败时返回 (0, 0)。
func ImageTtfbBox(size, angle float64, fontPath string, text string) (with int, heght int) {
	fontx, err := gntext.NewFont(fontPath)
	if err != nil {
		return 0, 0
	}
	fontx.SetDPI(72)
	fontx.SetPointSize(size)
	fontx.SetHinting(font.HintingNone)
	return fontx.MeasureText(text)
}

// GetFontFamily 从 fontPath 指定的 TTF 文件读取并解析字体。
// 文件读取或字体解析失败时返回错误；建议使用中文字体以避免中文变成方块。
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

// AutoWrap 按 width 与 line 自动将 text 拆分为多行。
//
// 参数依次为：字体大小、旋转角度、字体路径、文本、预设宽度、最大行数。
// 超出最大行数时会截断并追加 "......"。返回拆分后的字符串切片与最后一行的高度。
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

// LoadFontFace 从 TTF 文件 path 加载指定磅值的 font.Face，使用全 hinting。
// 读取或解析字体失败时返回错误。
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

// ImageJPEG 将图片 m 以 JPEG 编码写入 fpath。
// quantity 为图像质量（1-100）。创建文件或编码失败时返回错误，成功返回写入的文件路径。
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
	if err = jpeg.Encode(f, m, &jpeg.Options{Quality: quantity}); err != nil { // 写入文件
		fmt.Println("ImageJPEG:", err.Error())
		return "", err
	}

	file = fpath
	return
}

// ImagePNG 将图片 m 以 PNG 编码写入 fpath。
// 创建文件或编码失败时返回错误，成功返回写入的文件路径。
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
	if err = png.Encode(f, m); err != nil { // 写入文件
		fmt.Println("ImagePNG:", err.Error())
		return "", err
	}

	file = fpath
	return
}

// ImageGIF 将图片 m 以 GIF 编码写入 fpath。
// 创建文件或编码失败时返回错误，成功返回写入的文件路径。
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
	if err = gif.Encode(f, m, &gif.Options{}); err != nil { // 写入文件
		fmt.Println("ImageGIF:", err.Error())
		return "", err
	}

	file = fpath
	return
}

// ImageBMP 将图片 m 以 BMP 编码写入 fpath。
// 创建文件或编码失败时返回错误，成功返回写入的文件路径。
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
	if err = bmp.Encode(f, m); err != nil { // 写入文件
		fmt.Println("ImageBMP:", err.Error())
		return "", err
	}

	file = fpath
	return
}

// DrawText 在 target 上以指定字体、颜色与位置绘制单行文本。
// 文本绘制失败时返回错误，否则返回原图。
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
	pt := freetype.Pt(x, y)
	// 根据 Pt 的坐标值绘制给定的文本内容
	_, err := f.DrawString(text, pt)
	if err != nil {
		fmt.Println("DrawText:", err.Error())
		return nil, err
	}
	return target, nil
}

// DrawTextMultiline 在 target 上按行绘制多行文本，每行垂直间距为 height。
// 任何一行绘制失败会立即返回错误。
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
