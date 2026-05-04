package xio

import (
	"net/http"
	"os"
)

// GetFileTypeNameUsingMime 使用文件内容获取文件类型
func GetFileTypeNameUsingMime(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return GetFileTypeUsingMime(file)
}

// GetFileTypeUsingMime 使用文件内容获取文件类型
func GetFileTypeUsingMime(file *os.File) (string, error) {
	// 获取文件的 mimetype
	_, err := file.Stat()
	if err != nil {
		return "", err
	}
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	mimetype := http.DetectContentType(buffer)

	return mimetype, nil
}
