package xio

import (
	"context"
	"fmt"
	"os"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

//go:generate goption -p . -c PrintToPDFParams -w
//go:generate gofmt -w .

// PrintToPDFParams 打印参数
type PrintToPDFParams struct {
	Landscape               bool    `json:"landscape,omitempty"`               // 横向打印. 默认false.
	DisplayHeaderFooter     bool    `json:"displayHeaderFooter,omitempty"`     // 打印header和footer. 默认false.
	PrintBackground         bool    `json:"printBackground,omitempty"`         // 打印背景图.  默认false.
	Scale                   float64 `json:"scale,omitempty"`                   // 放缩因子. 默认为1.
	PaperWidth              float64 `json:"paperWidth,omitempty"`              // 页面宽度(英寸). 默认8.5英寸（美国Letter标准尺寸，和A4纸差不太多）.
	PaperHeight             float64 `json:"paperHeight,omitempty"`             // 页面高度(英寸). 默认11英寸(Letter标准尺寸).
	MarginTop               float64 `json:"marginTop"`                         // 上边距(英寸). 默认1cm (大约0.4 英寸).
	MarginBottom            float64 `json:"marginBottom"`                      // 底边距(英寸). 默认1cm (大约0.4 英寸).
	MarginLeft              float64 `json:"marginLeft"`                        // 左边距(英寸). 默认1cm (大约0.4 英寸).
	MarginRight             float64 `json:"marginRight"`                       // 右边距(英寸). 默认1cm (大约0.4 英寸).
	PageRanges              string  `json:"pageRanges,omitempty"`              // 要打印的页码, 比如, '1-5, 8, 11-13'.默认为空，全打印.
	IgnoreInvalidPageRanges bool    `json:"ignoreInvalidPageRanges,omitempty"` // 是否要忽略非法的页码范围. 默认false.
	HeaderTemplate          string  `json:"headerTemplate,omitempty"`          // HTML模板head.
	FooterTemplate          string  `json:"footerTemplate,omitempty"`          // HTML模板footer.
	PreferCSSPageSize       bool    `json:"preferCSSPageSize,omitempty"`       // 是否首选css定义的页面大小？默认false,将自动适应.
	// TransferMode            PrintToPDFTransferMode `json:"transferMode,omitempty"`            // 返回stream
}

// ChromedpPrintPdf
/**
 * @description: Chromedp打印pdf
 * @param {string} url
 * @param {string} to
 * @return {error}
 */
func ChromedpPrintPdf(url string, to string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				WithDisplayHeaderFooter(false).
				WithLandscape(false).
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return fmt.Errorf("chromedp Run failed,err:%+v", err)
	}
	if err := os.WriteFile(to, buf, 0o644); err != nil {
		return fmt.Errorf("write to file failed,err:%+v", err)
	}
	return nil
}

// GoWkhtmlPrintPdf
/**
 * @description: wkhtmltopdf打印pdf
 * @param {string} url
 * @param {string} to
 * @return {*}
 * @see https://wkhtmltopdf.org/downloads.html
 */
func GoWkhtmlPrintPdf(url string, to string) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println("GoWkhtmlPrintPdf---001>", err.Error())
		return err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage(url)

	// Set options for this page
	page.FooterRight.Set("[page]")
	// page.FooterFontSize.Set(10)
	page.Zoom.Set(1)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		fmt.Println("GoWkhtmlPrintPdf---002>", err.Error())
		fmt.Println(err)
		return err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(to)
	if err != nil {
		fmt.Println("GoWkhtmlPrintPdf---003>", err.Error())
		fmt.Println(err)
	}
	return err
}
