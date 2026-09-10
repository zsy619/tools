package xio

// PrintToPDFParamsOption 是一个用于配置 PrintToPDFParams 的可选参数函数类型。
type PrintToPDFParamsOption func(*PrintToPDFParams)

// NewPrintToPDFParams 创建一个新的 PrintToPDFParams 实例，并可传入若干选项函数进行初始化配置。
func NewPrintToPDFParams(opts ...PrintToPDFParamsOption) (printtopdfparams *PrintToPDFParams) {
	printtopdfparams = &PrintToPDFParams{}
	for _, opt := range opts {
		opt(printtopdfparams)
	}
	return
}

// WithPrintToPDFParamsLandscape 设置是否横向打印。
func WithPrintToPDFParamsLandscape(landscape bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.Landscape = landscape
	}
}

// WithPrintToPDFParamsDisplayHeaderFooter 设置是否打印页眉和页脚。
func WithPrintToPDFParamsDisplayHeaderFooter(displayheaderfooter bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.DisplayHeaderFooter = displayheaderfooter
	}
}

// WithPrintToPDFParamsPrintBackground 设置是否打印背景。
func WithPrintToPDFParamsPrintBackground(printbackground bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PrintBackground = printbackground
	}
}

// WithPrintToPDFParamsScale 设置页面缩放比例。
func WithPrintToPDFParamsScale(scale float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.Scale = scale
	}
}

// WithPrintToPDFParamsPaperWidth 设置页面宽度(英寸)。
func WithPrintToPDFParamsPaperWidth(paperwidth float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PaperWidth = paperwidth
	}
}

// WithPrintToPDFParamsPaperHeight 设置页面高度(英寸)。
func WithPrintToPDFParamsPaperHeight(paperheight float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PaperHeight = paperheight
	}
}

// WithPrintToPDFParamsMarginTop 设置上边距(英寸)。
func WithPrintToPDFParamsMarginTop(margintop float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginTop = margintop
	}
}

// WithPrintToPDFParamsMarginBottom 设置下边距(英寸)。
func WithPrintToPDFParamsMarginBottom(marginbottom float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginBottom = marginbottom
	}
}

// WithPrintToPDFParamsMarginLeft 设置左边距(英寸)。
func WithPrintToPDFParamsMarginLeft(marginleft float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginLeft = marginleft
	}
}

// WithPrintToPDFParamsMarginRight 设置右边距(英寸)。
func WithPrintToPDFParamsMarginRight(marginright float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginRight = marginright
	}
}

// WithPrintToPDFParamsPageRanges 设置要打印的页码范围。
func WithPrintToPDFParamsPageRanges(pageranges string) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PageRanges = pageranges
	}
}

// WithPrintToPDFParamsIgnoreInvalidPageRanges 设置是否忽略非法的页码范围。
func WithPrintToPDFParamsIgnoreInvalidPageRanges(ignoreinvalidpageranges bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.IgnoreInvalidPageRanges = ignoreinvalidpageranges
	}
}

// WithPrintToPDFParamsHeaderTemplate 设置页眉的 HTML 模板。
func WithPrintToPDFParamsHeaderTemplate(headertemplate string) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.HeaderTemplate = headertemplate
	}
}

// WithPrintToPDFParamsFooterTemplate 设置页脚的 HTML 模板。
func WithPrintToPDFParamsFooterTemplate(footertemplate string) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.FooterTemplate = footertemplate
	}
}

// WithPrintToPDFParamsPreferCSSPageSize 设置是否优先采用 CSS 定义的页面大小。
func WithPrintToPDFParamsPreferCSSPageSize(prefercsspagesize bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PreferCSSPageSize = prefercsspagesize
	}
}
