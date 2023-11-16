package xio

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Zip
/**
 * @description: 压缩文件
 * @param {string} src_dir 源文件夹
 * @param {string} zip_file_name 压缩文件名
 * @param {bool} zip_header_name 是否压缩文件夹
 * @return {error}
 */
func Zip(src_dir string, zip_file_name string, zip_header_name bool) error {
	dir, err := os.ReadDir(src_dir)
	if err != nil {
		return err
	}
	if len(dir) == 0 {
		return nil
	}
	// 预防：旧文件无法覆盖
	os.RemoveAll(zip_file_name)
	// 创建：zip文件
	zipfile, _ := os.Create(zip_file_name)
	defer zipfile.Close()
	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	// 遍历路径信息
	filepath.Walk(src_dir, func(path string, info os.FileInfo, _ error) error {
		// 如果是源路径，提前进行下一个遍历
		if path == src_dir {
			return nil
		}
		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		if zip_header_name {
			header.Name = strings.TrimPrefix(path, src_dir+`\`)
		}
		// 判断：文件是不是文件夹
		if info.IsDir() {
			if zip_header_name {
				header.Name += `/`
			}
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}
		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
	return nil
}
