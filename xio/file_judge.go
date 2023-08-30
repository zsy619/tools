package xio

import (
	"path"
	"strings"
)

// IsImage 判断文件是否为图片
// ".bmp", ".png", ".jpg", ".jpeg", ".gif", ".webp"
func IsImage(filename string) bool {
	// 判断文件是否为图片
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".bmp", ".png", ".jpg", ".jpeg", ".gif", ".webp":
		return true
	}
	return false
}

// IsScript 判断文件是否为脚本
func IsScript(filename string) bool {
	// 判断文件是否为图片
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".js", ".php", ".sh", ".shell", ".py", ".rs":
		return true
	}
	return false
}

// IsAllowFile 判断文件是否为允许的文件
// ".bmp", ".png", ".jpg", ".jpeg", ".gif", ".webp", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx", ".doc", ".docx"
func IsAllowFile(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".bmp", ".png", ".jpg", ".jpeg", ".webp", ".gif", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx", ".doc", ".docx":
		return true
	}
	return false
}

// IsAllowExcel 判断文件是否为允许的excel文件
// ".xls", ".xlsx"
func IsAllowExcelXlsx(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".xlsx":
		return true
	}
	return false
}

// IsAllowExcel 判断文件是否为允许的excel文件
// ".xls", ".xlsx"
func IsAllowExcel(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".xls", ".xlsx":
		return true
	}
	return false
}

// IsCompress 判断文件是否为压缩文件
// ".zip", ".rar", ".tar", ".gz", ".7z"
func IsCompress(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".zip", ".rar", ".tar", ".gz", ".7z":
		return true
	}
	return false
}

// IsFont 判断文件是否为字体文件
// ".ttf", ".otf", ".woff", ".woff2"
func IsFont(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".ttf", ".otf", ".woff", ".woff2":
		return true
	}
	return false
}

// IsVideo 判断文件是否为视频文件
// ".mp3", ".wav", ".ogg", ".m4a"
func IsAudio(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".mp3", ".wav", ".ogg", ".m4a":
		return true
	}
	return false
}

// IsVideo 判断文件是否为视频文件
// ".mp4", ".avi", ".mkv", ".rmvb", ".rm", ".flv", ".mov", ".wmv", ".3gp", ".mpeg", ".mpg"
func IsVideo(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".mp4", ".avi", ".mkv", ".rmvb", ".rm", ".flv", ".mov", ".wmv", ".3gp", ".mpeg", ".mpg":
		return true
	}
	return false
}

// IsText 判断文件是否为文本文件
// ".txt", ".md", ".markdown", ".json", ".xml", ".html", ".htm", ".css"
func IsText(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".txt", ".md", ".markdown", ".json", ".xml", ".html", ".htm", ".css":
		return true
	}
	return false
}

// IsPDF 判断文件是否为PDF文件
// ".pdf"
func IsPDF(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".pdf":
		return true
	}
	return false
}

// IsWord 判断文件是否为Word文件
// ".doc", ".docx"
func IsWord(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".doc", ".docx":
		return true
	}
	return false
}

// IsPPT 判断文件是否为PPT文件
// ".ppt", ".pptx"
func IsPPT(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".ppt", ".pptx":
		return true
	}
	return false
}
