package xio

import (
	"path"
	"strings"
)

// IsImage 判断文件是否为图片
func IsImage(filename string) bool {
	// 判断文件是否为图片
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".bmp", ".png", ".jpg", ".jpeg", ".gif":
		isImage = true
	}
	return isImage
}

// IsScript 判断文件是否为脚本
func IsScript(filename string) bool {
	// 判断文件是否为图片
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".js", ".php", ".sh", ".shell", ".py", ".rs":
		isImage = true
	}
	return isImage
}

// IsAllowFile 判断文件是否为允许的文件
// ".bmp", ".png", ".jpg", ".jpeg", ".gif", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx", ".doc", ".docx"
func IsAllowFile(filename string) bool {
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".bmp", ".png", ".jpg", ".jpeg", ".gif", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx", ".doc", ".docx":
		isImage = true
	}
	return isImage
}

// IsCompress 判断文件是否为压缩文件
func IsCompress(filename string) bool {
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".zip", ".rar", ".tar", ".gz", ".7z":
		isImage = true
	}
	return isImage
}

// IsFont 判断文件是否为字体文件
func IsFont(filename string) bool {
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".ttf", ".otf", ".woff", ".woff2":
		isImage = true
	}
	return isImage
}

// IsVideo 判断文件是否为视频文件
func IsAudio(filename string) bool {
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".mp3", ".wav", ".ogg", ".m4a":
		isImage = true
	}
	return isImage
}

// IsVideo 判断文件是否为视频文件
func IsVideo(filename string) bool {
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".mp4", ".avi", ".mkv", ".rmvb", ".rm", ".flv", ".mov", ".wmv", ".3gp", ".mpeg", ".mpg":
		isImage = true
	}
	return isImage
}

// IsText 判断文件是否为文本文件
func IsText(filename string) bool {
	isImage := false
	ext := path.Ext(filename)
	ext = strings.ToLower(ext)
	switch ext {
	case ".txt", ".md", ".markdown", ".json", ".xml", ".html", ".htm", ".css":
		isImage = true
	}
	return isImage
}
