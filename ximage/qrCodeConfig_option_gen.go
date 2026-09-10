package ximage

import qrcode "github.com/skip2/go-qrcode"

// QrCodeConfigOption 是用于构造 *QrCodeConfig 的函数式选项。
type QrCodeConfigOption func(*QrCodeConfig)

// NewQrCodeConfig 根据给定的选项构造一个 QrCodeConfig。
// 不传选项时返回零值结构体。
func NewQrCodeConfig(opts ...QrCodeConfigOption) (qrcodeconfig *QrCodeConfig) {
	qrcodeconfig = &QrCodeConfig{}
	for _, opt := range opts {
		opt(qrcodeconfig)
	}
	return
}

// WithQrCodeConfigLogoPath 设置 QrCodeConfig.LogoPath 的函数式选项。
func WithQrCodeConfigLogoPath(logopath string) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.LogoPath = logopath
	}
}

// WithQrCodeConfigLogoResize 设置 QrCodeConfig.LogoResize 的函数式选项。
func WithQrCodeConfigLogoResize(logoresize Resize) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.LogoResize = logoresize
	}
}

// WithQrCodeConfigContent 设置 QrCodeConfig.Content 的函数式选项。
func WithQrCodeConfigContent(content string) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Content = content
	}
}

// WithQrCodeConfigLevel 设置 QrCodeConfig.Level 的函数式选项。
func WithQrCodeConfigLevel(level qrcode.RecoveryLevel) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Level = level
	}
}

// WithQrCodeConfigSize 设置 QrCodeConfig.Size 的函数式选项。
func WithQrCodeConfigSize(size int) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Size = size
	}
}

// WithQrCodeConfigOffset 设置 QrCodeConfig.Offset 的函数式选项。
func WithQrCodeConfigOffset(offset Offset) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Offset = offset
	}
}
