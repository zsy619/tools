package ximage

// TitleConfigOption 是用于构造 *TitleConfig 的函数式选项。
type TitleConfigOption func(*TitleConfig)

// NewTitleConfig 根据给定的选项构造一个 TitleConfig。
// 不传选项时返回零值结构体。
func NewTitleConfig(opts ...TitleConfigOption) (titleconfig *TitleConfig) {
	titleconfig = &TitleConfig{}
	for _, opt := range opts {
		opt(titleconfig)
	}
	return
}

// WithTitleConfigTitle 设置 TitleConfig.Title 的函数式选项。
func WithTitleConfigTitle(title string) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Title = title
	}
}

// WithTitleConfigOffset 设置 TitleConfig.Offset 的函数式选项。
func WithTitleConfigOffset(offset Offset) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Offset = offset
	}
}

// WithTitleConfigHeight 设置 TitleConfig.Height 的函数式选项。
func WithTitleConfigHeight(height int) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Height = height
	}
}

// WithTitleConfigHinting 设置 TitleConfig.Hinting 的函数式选项。
func WithTitleConfigHinting(hinting string) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Hinting = hinting
	}
}

// WithTitleConfigDpi 设置 TitleConfig.Dpi 的函数式选项。
func WithTitleConfigDpi(dpi float64) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Dpi = dpi
	}
}

// WithTitleConfigSize 设置 TitleConfig.Size 的函数式选项。
func WithTitleConfigSize(size float64) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Size = size
	}
}

// WithTitleConfigSpacing 设置 TitleConfig.Spacing 的函数式选项。
func WithTitleConfigSpacing(spacing float64) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Spacing = spacing
	}
}

// WithTitleConfigPath 设置 TitleConfig.Path 的函数式选项。
func WithTitleConfigPath(path string) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Path = path
	}
}

// WithTitleConfigWonb 设置 TitleConfig.Wonb 的函数式选项。
func WithTitleConfigWonb(wonb bool) func(*TitleConfig) {
	return func(titleconfig *TitleConfig) {
		titleconfig.Wonb = wonb
	}
}
