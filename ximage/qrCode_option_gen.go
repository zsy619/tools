package ximage

// QrCodeOption 是用于构造 *QrCode 的函数式选项。
type QrCodeOption func(*QrCode)

// NewQrCode 根据给定的选项构造一个 QrCode。
// 不传选项时返回零值结构体。
func NewQrCode(opts ...QrCodeOption) (qrcode *QrCode) {
	qrcode = &QrCode{}
	for _, opt := range opts {
		opt(qrcode)
	}
	return
}

// WithQrCodeIsOne 设置 QrCode.IsOne 的函数式选项：是否按内容哈希复用生成一次。
func WithQrCodeIsOne(isone bool) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.IsOne = isone
	}
}

// WithQrCodeRoot 设置 QrCode.Root（最终图片保存路径）的函数式选项。
func WithQrCodeRoot(root string) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.Root = root
	}
}

// WithQrCodeTitle 设置 QrCode.Title 的函数式选项。
func WithQrCodeTitle(title TitleConfig) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.Title = title
	}
}

// WithQrCodeBackground 设置 QrCode.Background 的函数式选项。
func WithQrCodeBackground(background BackgroundConfig) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.Background = background
	}
}

// WithQrCodeQrCodeConfig 设置 QrCode.QrCodeConfig 的函数式选项。
func WithQrCodeQrCodeConfig(qrcodeconfig QrCodeConfig) func(*QrCode) {
	return func(qrcode *QrCode) {
		qrcode.QrCodeConfig = qrcodeconfig
	}
}
