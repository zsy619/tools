package ximage

import "testing"

func TestQrCode_CreateQrCodeBackground(t *testing.T) {
	type fields struct {
		Root       string
		Title      TitleConfig
		Background BackgroundConfig
		QrCode     QrCodeConfig
	}
	tests := []struct {
		name     string
		fields   fields
		wantFile string
		wantErr  bool
	}{
		// TODO: Add test cases.
		// {"t1", fields{
		// 	Title: TitleConfig{},
		// 	Background: BackgroundConfig{
		// 		Path: "../../static/public/Home/image/qrcode_bottom.png",
		// 	},
		// 	QrCode: QrCodeConfig{
		// 		Content: "合同签订后 21 个工作日内支付本合同总金额的 50% ，即人民币 15 万元整（大写：人民币拾伍万元整）",
		// 		Offset:  Offset{Y: 100, X: 0},
		// 	},
		// }, "", false},
		// {"t2", fields{
		// 	Title: TitleConfig{},
		// 	Background: BackgroundConfig{
		// 		Path: "../../static/public/Home/image/qrcode_bottom.png",
		// 	},
		// 	QrCode: QrCodeConfig{
		// 		Content:  "合同签订后 21 个工作日内支付本合同总金额的 50% ，即人民币 15 万元整（大写：人民币拾伍万元整）",
		// 		Offset:   Offset{Y: 100, X: 0},
		// 		LogoPath: "../../static/public/Admin/image/show_logo.png",
		// 	},
		// }, "", false},
		{"t3", fields{
			Title: TitleConfig{
				Title:  "河南省教育网有限公司",
				Path:   "../../static/fonts/simhei.ttf",
				Size:   48,
				Dpi:    72,
				Offset: Offset{Y: 370, X: 0},
			},
			Background: BackgroundConfig{
				Path: "../../static/public/Home/image/qrcode_bottom.png",
			},
			QrCode: QrCodeConfig{
				Content: "合同签订后 21 个工作日内支付本合同总金额的 50% ，即人民币 15 万元整（大写：人民币拾伍万元整）",
				Offset:  Offset{Y: 100, X: 0},
				// LogoPath: "../../static/public/Admin/image/show_logo.png",
			},
		}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &QrCode{
				Root:         tt.fields.Root,
				Title:        tt.fields.Title,
				Background:   tt.fields.Background,
				QrCodeConfig: tt.fields.QrCode,
			}
			gotFile, err := this.CreateQrCodeBackground()
			if (err != nil) != tt.wantErr {
				t.Errorf("QrCode.CreateQrCodeBackground() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFile != "" {
				t.Logf("QrCode.CreateQrCodeBackground() = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
