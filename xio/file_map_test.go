package xio

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFileType(t *testing.T) {
	f, err := os.Open("/Users/zhushuyan/Downloads/WechatIMG3264.jpg")
	if err != nil {
		t.Logf("open error: %v", err)
	}

	defer f.Close()

	fSrc, err := ioutil.ReadAll(f)
	t.Log(GetFileType(fSrc[:10]))
}
