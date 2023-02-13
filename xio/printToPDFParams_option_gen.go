package xio

type PrintToPDFParamsOption func(*PrintToPDFParams)

func NewPrintToPDFParams(opts ...PrintToPDFParamsOption) (printtopdfparams *PrintToPDFParams) {
	printtopdfparams = &PrintToPDFParams{}
	for _, opt := range opts {
		opt(printtopdfparams)
	}
	return
}

func WithPrintToPDFParamsLandscape(landscape bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.Landscape = landscape
	}
}

func WithPrintToPDFParamsDisplayHeaderFooter(displayheaderfooter bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.DisplayHeaderFooter = displayheaderfooter
	}
}

func WithPrintToPDFParamsPrintBackground(printbackground bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PrintBackground = printbackground
	}
}

func WithPrintToPDFParamsScale(scale float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.Scale = scale
	}
}

func WithPrintToPDFParamsPaperWidth(paperwidth float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PaperWidth = paperwidth
	}
}

func WithPrintToPDFParamsPaperHeight(paperheight float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PaperHeight = paperheight
	}
}

func WithPrintToPDFParamsMarginTop(margintop float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginTop = margintop
	}
}

func WithPrintToPDFParamsMarginBottom(marginbottom float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginBottom = marginbottom
	}
}

func WithPrintToPDFParamsMarginLeft(marginleft float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginLeft = marginleft
	}
}

func WithPrintToPDFParamsMarginRight(marginright float64) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.MarginRight = marginright
	}
}

func WithPrintToPDFParamsPageRanges(pageranges string) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PageRanges = pageranges
	}
}

func WithPrintToPDFParamsIgnoreInvalidPageRanges(ignoreinvalidpageranges bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.IgnoreInvalidPageRanges = ignoreinvalidpageranges
	}
}

func WithPrintToPDFParamsHeaderTemplate(headertemplate string) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.HeaderTemplate = headertemplate
	}
}

func WithPrintToPDFParamsFooterTemplate(footertemplate string) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.FooterTemplate = footertemplate
	}
}

func WithPrintToPDFParamsPreferCSSPageSize(prefercsspagesize bool) func(*PrintToPDFParams) {
	return func(printtopdfparams *PrintToPDFParams) {
		printtopdfparams.PreferCSSPageSize = prefercsspagesize
	}
}
