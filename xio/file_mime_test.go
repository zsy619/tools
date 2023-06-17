package xio

import (
	"os"
	"testing"
)

func TestGetFileTypeUsingMime(t *testing.T) {
	file, err := os.Open("/Users/zhushuyan/Downloads/WechatIMG67.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	mimetype, err := GetFileTypeUsingMime(file)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mimetype)
}

func TestGetFileTypeNameUsingMime(t *testing.T) {
	fileNames := []string{
		// "/Users/zhushuyan/Downloads/WechatIMG67.jpeg",
		// "/Users/zhushuyan/Downloads/webOS.zip",
		// "/Users/zhushuyan/Downloads/jh.bat",
		// "/Users/zhushuyan/Downloads/PMBOK第六版-中文版.pdf",
		// "/Users/zhushuyan/Downloads/BingSiteAuth.xml",
		// "/Users/zhushuyan/Downloads/运维平台新增需求点梳理-20230602.xlsx",
		// "/Users/zhushuyan/Downloads/超级5D全景S系列安装使用说明.docx",
		// "/Users/zhushuyan/Downloads/修改.doc",
		"/Users/zhushuyan/Downloads/帮助文档/images/logo.svg",
		"/Users/zhushuyan/Downloads/帮助文档/js/jquery-1.7.2.min.js",
		"/Users/zhushuyan/Downloads/帮助文档/index.html",
		"/Users/zhushuyan/Downloads/帮助文档/css/main.css",
		"/Users/zhushuyan/Downloads/帮助文档/css/main.less",
		"/Users/zhushuyan/Downloads/帮助文档/fonts/fontawesome-webfont.eot",
		"/Users/zhushuyan/Downloads/帮助文档/fonts/fontawesome-webfont.ttf",
		"/Users/zhushuyan/Downloads/帮助文档/fonts/fontawesome-webfont.woff",
		"/Users/zhushuyan/Downloads/帮助文档/fonts/fontawesome-webfont.woff2",
		"/Users/zhushuyan/Downloads/帮助文档/fonts/FontAwesome.otf",
	}

	for _, fileName := range fileNames {
		mimetype, err := GetFileTypeNameUsingMime(fileName)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(fileName, mimetype)
	}
}
