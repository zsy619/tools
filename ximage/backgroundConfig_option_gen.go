package ximage

type BackgroundConfigOption func(*BackgroundConfig)

func NewBackgroundConfig(opts ...BackgroundConfigOption) (backgroundconfig *BackgroundConfig) {
	backgroundconfig = &BackgroundConfig{}
	for _, opt := range opts {
		opt(backgroundconfig)
	}
	return
}

func WithBackgroundConfigPath(path string) func(*BackgroundConfig) {
	return func(backgroundconfig *BackgroundConfig) {
		backgroundconfig.Path = path
	}
}

func WithBackgroundConfigResize(resize Resize) func(*BackgroundConfig) {
	return func(backgroundconfig *BackgroundConfig) {
		backgroundconfig.Resize = resize
	}
}
