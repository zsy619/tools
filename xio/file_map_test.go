package xio

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/h2non/filetype"
)

func TestGetFileType(t *testing.T) {
	f, err := os.Open("/Users/zhushuyan/Downloads/WechatIMG67.jpeg")
	if err != nil {
		t.Logf("open error: %v", err)
	}

	defer f.Close()

	fSrc, err := io.ReadAll(f)
	if err != nil {
		t.Logf("read error: %v", err)
	} else {
		t.Log(GetFileType(fSrc[:10]))
	}
}

func TestGetFileTypeByFileName(t *testing.T) {
	fileNames := []string{
		"/Users/zhushuyan/Downloads/WechatIMG67.jpeg",
		"/Users/zhushuyan/Downloads/webOS.zip",
		"/Users/zhushuyan/Downloads/jh.bat",
		"/Users/zhushuyan/Downloads/PMBOK第六版-中文版.pdf",
		"/Users/zhushuyan/Downloads/BingSiteAuth.xml",
		"/Users/zhushuyan/Downloads/运维平台新增需求点梳理-20230602.xlsx",
		"/Users/zhushuyan/Downloads/超级5D全景S系列安装使用说明.docx",
		"/Users/zhushuyan/Downloads/修改.doc",
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
		fileType, err := GetFileTypeByFileName(fileName)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Log(fileName, fileType)
		}
	}
}

func TestXxx(t *testing.T) {
	// 打开文件
	f, err := os.Open("/Users/zhushuyan/Downloads/帮助文档/images/logo.svg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 获取文件类型
	kind, _ := filetype.MatchReader(f)
	fmt.Printf("File type: %s\n", kind.MIME.Value)
}
