package ximage

import qrcode "github.com/skip2/go-qrcode"

type QrCodeConfigOption func(*QrCodeConfig)

func NewQrCodeConfig(opts ...QrCodeConfigOption) (qrcodeconfig *QrCodeConfig) {
	qrcodeconfig = &QrCodeConfig{}
	for _, opt := range opts {
		opt(qrcodeconfig)
	}
	return
}

func WithQrCodeConfigLogoPath(logopath string) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.LogoPath = logopath
	}
}

func WithQrCodeConfigLogoResize(logoresize Resize) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.LogoResize = logoresize
	}
}

func WithQrCodeConfigContent(content string) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Content = content
	}
}

func WithQrCodeConfigLevel(level qrcode.RecoveryLevel) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Level = level
	}
}

func WithQrCodeConfigSize(size int) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Size = size
	}
}

func WithQrCodeConfigOffset(offset Offset) func(*QrCodeConfig) {
	return func(qrcodeconfig *QrCodeConfig) {
		qrcodeconfig.Offset = offset
	}
}
