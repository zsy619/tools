package ximage

type TitleConfigOption func(*TitleConfig)

func NewTitleConfig(opts ...TitleConfigOption) (titleconfig *TitleConfig) {
	titleconfig = &TitleConfig{}
	for _, opt := range opts {
		opt(titleconfig)
	}
	return
}

func WithTitleConfigTitle(title string) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Title = title
	}
}

func WithTitleConfigOffset(offset Offset) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Offset = offset
	}
}

func WithTitleConfigHeight(height int) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Height = height
	}
}

func WithTitleConfigHinting(hinting string) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Hinting = hinting
	}
}

func WithTitleConfigDpi(dpi float64) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Dpi = dpi
	}
}

func WithTitleConfigSize(size float64) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Size = size
	}
}

func WithTitleConfigSpacing(spacing float64) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Spacing = spacing
	}
}

func WithTitleConfigPath(path string) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Path = path
	}
}

func WithTitleConfigWonb(wonb bool) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Wonb = wonb
	}
}
