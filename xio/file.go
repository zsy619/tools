package xio

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	COMMON_FILE_MODE = 0o644 // common file mode , value as '-rw-r--r--' by unix os
)

// file info watch call back function process during in #ListFiles
type callbackFn func(filename, path string, data []byte)

func ListFiles(dir string, callback callbackFn) (reterr error) {
	fileSystem := os.DirFS(dir)
	return fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := os.ReadFile(filepath.Join(dir, path))
		if err != nil {
			return err
		}
		callback(path, dir, data)
		return nil
	})
}

// IsDir 存在且为目录
func IsDir(dir string) bool {
	fi, err := os.Stat(dir)
	if err == nil {
		return fi.IsDir()
	}
	return false
}

// IsFile 存在且为文件
func IsFile(file string) bool {
	fi, err := os.Stat(file)
	if err == nil {
		return !fi.IsDir()
	}
	return false
}

/* 获取指定路径下以及所有子目录下的所有文件，可匹配后缀过滤（suffix为空则不过滤）*/
func WalkDir(dir, suffix string) (files []string, err error) {
	files = []string{}
	err = filepath.Walk(dir, func(fname string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			// 忽略目录
			return nil
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			// 文件后缀匹配
			files = append(files, fname)
		}
		return nil
	})

	return files, err
}

/* 获取指定路径下的所有文件，只搜索当前路径，不进入下一级目录，可匹配后缀过滤（suffix为空则不过滤）*/
func ListDir(dir, suffix string) (files []string, err error) {
	files = []string{}
	_dir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToLower(suffix) // 匹配后缀
	for _, _file := range _dir {
		if _file.IsDir() {
			continue // 忽略目录
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
			// 文件后缀匹配
			files = append(files, path.Join(dir, _file.Name()))
		}
	}
	return files, nil
}

// 拷贝挣个文件夹
func CopyDir(srcPath, desPath string) error {
	// 检查目录是否正确
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("源路径与目标路径不能相同！")
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		// 复制目录是将源目录中的子目录复制到目标路径中，不包含源目录本身
		if path == srcPath {
			return nil
		}
		// 生成新路径
		destNewPath := strings.Replace(path, srcPath, desPath, -1)
		if !f.IsDir() {
			CopyFile(path, destNewPath)
		} else {
			if !FileIsExisted(destNewPath) {
				return MakeDir(destNewPath)
			}
		}
		return nil
	})

	return err
}

// 使用io.Copy
func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	// 获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	// desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) // 复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()
	return io.Copy(desFile, srcFile)
}

// 使用os.WriteFile()和os.ReadFile()
func CopyFile2(src, des string) (written int64, err error) {
	// 获取源文件的权限
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	srcFile.Close()

	input, err := os.ReadFile(src)
	if err != nil {
		return 0, err
	}
	err = os.WriteFile(des, input, perm)
	if err != nil {
		return 0, err
	}
	return int64(len(input)), nil
}

// 使用os.Read()和os.Write()
func CopyFile3(src, des string, bufSize int) (written int64, err error) {
	if bufSize <= 0 {
		bufSize = 1 * 1024 * 1024 // 1M
	}
	buf := make([]byte, bufSize)
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	// 获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	desFile, err := os.OpenFile(des, os.O_CREATE|os.O_RDWR|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer desFile.Close()
	count := 0
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			break
		}
		if wn, err := desFile.Write(buf[:n]); err != nil {
			return 0, err
		} else {
			count += wn
		}
	}
	return int64(count), nil
}

// 创建文件夹（如果文件夹不存在则创建）
func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0o777); err != nil { // os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

// 检查文件是否存在
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}
