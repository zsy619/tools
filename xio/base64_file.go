package xio

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

// 写入文件,保存
func Base64ToFile(path string, base64_image_content string, datePath bool) (string, error) {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64_image_content)
	if !b {
		return "", fmt.Errorf("base64 image content is not valid")
	}

	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(base64_image_content), 2)
	fileType := string(allData[0][1]) // png ，jpeg 后缀获取

	base64Str := re.ReplaceAllString(base64_image_content, "")

	if datePath {
		date := time.Now().Format("2006-01-02")
		if ok := IsFileExist(path + "/" + date); !ok {
			os.Mkdir(path+"/"+date, 0o666)
		}
		path = path + "/" + date
	}

	curFileStr := strconv.FormatInt(time.Now().UnixNano(), 10)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(99999)

	var file string = path + "/" + curFileStr + strconv.Itoa(n) + "." + fileType
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(file, byte, 0o666)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return file, nil
}

// 判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
