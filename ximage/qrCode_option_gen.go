package ximage

type QrCodeOption func(*QrCode)

func NewQrCode(opts ...QrCodeOption) (qrcode *QrCode) {
	qrcode = &QrCode{}
	for _, opt := range opts {
		opt(qrcode)
	}
	return
}

func WithQrCodeIsOne(isone bool) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.IsOne = isone
	}
}

func WithQrCodeRoot(root string) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.Root = root
	}
}

func WithQrCodeTitle(title TitleConfig) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.Title = title
	}
}

func WithQrCodeBackground(background BackgroundConfig) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.Background = background
	}
}

func WithQrCodeQrCodeConfig(qrcodeconfig QrCodeConfig) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.QrCodeConfig = qrcodeconfig
	}
}
