package ximage

// BackgroundConfigOption 是用于构造 *BackgroundConfig 的函数式选项。
type BackgroundConfigOption func(*BackgroundConfig)

// NewBackgroundConfig 根据给定的选项构造一个 BackgroundConfig。
// 不传选项时返回零值结构体。
func NewBackgroundConfig(opts ...BackgroundConfigOption) (backgroundconfig *BackgroundConfig) {
	backgroundconfig = &BackgroundConfig{}
	for _, opt := range opts {
		opt(backgroundconfig)
	}
	return
}

// WithBackgroundConfigPath 设置 BackgroundConfig.Path 的函数式选项。
func WithBackgroundConfigPath(path string) func(*BackgroundConfig) {
	return func(backgroundconfig *BackgroundConfig) {
		backgroundconfig.Path = path
	}
}

// WithBackgroundConfigResize 设置 BackgroundConfig.Resize 的函数式选项。
func WithBackgroundConfigResize(resize Resize) func(*BackgroundConfig) {
	return func(backgroundconfig *BackgroundConfig) {
		backgroundconfig.Resize = resize
	}
}
