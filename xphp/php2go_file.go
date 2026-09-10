package xphp

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// Stat 实现 PHP stat()：返回 filename 的文件信息。
func Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}

// Pathinfo 实现 PHP pathinfo()。
//
// options 位标志：1=dirname、2=basename、4=extension、8=filename；-1 表示全部。
// 示例：Pathinfo("/home/go/path/src/php2go/php2go.go", 1|2|4|8)。
func Pathinfo(path string, options int) map[string]string {
	if options == -1 {
		options = 1 | 2 | 4 | 8
	}
	info := make(map[string]string)
	if (options & 1) == 1 {
		info["dirname"] = filepath.Dir(path)
	}
	if (options & 2) == 2 {
		info["basename"] = filepath.Base(path)
	}
	if ((options & 4) == 4) || ((options & 8) == 8) {
		basename := ""
		if (options & 2) == 2 {
			var ok bool
			basename, ok = info["basename"]
			if !ok {
				fmt.Println("basename not found")
			}
		} else {
			basename = filepath.Base(path)
		}
		p := strings.LastIndex(basename, ".")
		filename, extension := "", ""
		if p > 0 {
			filename, extension = basename[:p], basename[p+1:]
		} else if p == -1 {
			filename = basename
		} else if p == 0 {
			extension = basename[p+1:]
		}
		if (options & 4) == 4 {
			info["extension"] = extension
		}
		if (options & 8) == 8 {
			info["filename"] = filename
		}
	}
	return info
}

// FileExists 实现 PHP file_exists()：判断文件或目录是否存在。
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFile 实现 PHP is_file()：判断是否为常规文件。
// 注意：当前实现与 FileExists 等价，不严格区分目录。
func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir 实现 PHP is_dir()：判断是否为目录。
func IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// FileSize 实现 PHP filesize()：返回文件大小（字节）。
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// FilePutContents 实现 PHP file_put_contents()：以指定权限写入 data 到 filename。
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return os.WriteFile(filename, []byte(data), mode)
}

// FileGetContents 实现 PHP file_get_contents()：读取整个文件为字符串。
func FileGetContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

// Unlink 实现 PHP unlink()：删除文件。
func Unlink(filename string) error {
	return os.Remove(filename)
}

// Delete 是 Unlink 的别名。
func Delete(filename string) error {
	return os.Remove(filename)
}

// Copy 实现 PHP copy()：把 source 复制到 dest。
// 成功返回 (true, nil)；失败返回 (false, 错误)。
func Copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

// IsReadable 实现 PHP is_readable()：判断 filename 是否可读。
// 当前使用 syscall.Open + O_RDONLY 探测；权限不足时返回 false。
func IsReadable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_RDONLY, 0)
	return err == nil
}

// IsWriteable 实现 PHP is_writeable()：判断 filename 是否可写。
func IsWriteable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_WRONLY, 0)
	return err == nil
}

// Rename 实现 PHP rename()。
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// Touch 实现 PHP touch()：若文件不存在则创建，存在则不修改。
// 注意：当前不会更新已有文件的修改时间。
func Touch(filename string) (bool, error) {
	fd, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0o666)
	if err != nil {
		return false, err
	}
	fd.Close()
	return true, nil
}

// Mkdir 实现 PHP mkdir()：创建单层目录。
func Mkdir(filename string, mode os.FileMode) error {
	return os.Mkdir(filename, mode)
}

// Getcwd 实现 PHP getcwd()：返回当前工作目录。
func Getcwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

// Realpath 实现 PHP realpath()：返回绝对路径。
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// Basename 实现 PHP basename()：返回路径的最后一个元素。
func Basename(path string) string {
	return filepath.Base(path)
}

// Chmod 实现 PHP chmod()：修改文件权限；返回是否成功。
func Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// Chown 实现 PHP chown()：修改文件属主；返回是否成功。
func Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}

// Fclose 实现 PHP fclose()：关闭打开的文件。
func Fclose(handle *os.File) error {
	return handle.Close()
}

// Filemtime 实现 PHP filemtime()：返回文件的最后修改时间（Unix 秒）。
func Filemtime(filename string) (int64, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	fileinfo, err := fd.Stat()
	if err != nil {
		return 0, err
	}
	return fileinfo.ModTime().Unix(), nil
}

// Fgetcsv 实现 PHP fgetcsv()：从 handle 读取全部 CSV 记录，delimiter 为字段分隔符。
// 当前未使用 length 参数（TODO）。
func Fgetcsv(handle *os.File, length int, delimiter rune) ([][]string, error) {
	reader := csv.NewReader(handle)
	reader.Comma = delimiter
	// TODO length limit
	return reader.ReadAll()
}

// Glob 实现 PHP glob()：按 pattern 返回匹配的文件路径列表。
func Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}
