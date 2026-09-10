package xphp

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// IsNumeric 实现 PHP is_numeric() 的语义。
// 数值字符串由可选符号、若干数字、可选小数部分和可选指数部分组成，例如 +0123.45e6。
// 与 PHP 不同：本实现额外支持十六进制（0x 前缀）；空字符串返回 false。
func IsNumeric(val interface{}) bool {
	switch val := val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.TrimSpace(str)
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9, Point, Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

// Empty 实现 PHP empty() 的判定。
//
// 判定规则：
//   - nil（接口/指针）：视为空；
//   - 字符串/数组：长度为 0 视为空；
//   - Map/Slice：长度为 0 或为 nil 视为空；
//   - 布尔：值为 false 视为空；
//   - 整型/浮点：值为 0 视为空；
//   - 其它：与对应零值比较；
//   - 变量不存在时也不会产生告警。
func Empty(val interface{}) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// Exec 实现 PHP exec() 的语义。
//
// returnVar：0 表示成功，1 表示失败；output 接收按行切分的命令输出（去掉末尾换行）。
// 返回值为 output 的最后一行；若命令执行失败则返回空串。
// 命令字符串支持形如 "ls -a" 或 "/bin/bash -c \"ls -a\"" 的格式。
func Exec(command string, output *[]string, returnVar *int) string {
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})
	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}
	cmd := exec.Command(parts[0], parts[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		*returnVar = 1
		return ""
	}
	*returnVar = 0
	*output = strings.Split(strings.TrimRight(string(out), "\n"), "\n")
	if l := len(*output); l > 0 {
		return (*output)[l-1]
	}
	return ""
}

// System 实现 PHP system() 的语义：实时把 stdout/stderr 打印到当前进程并返回最后一行。
// returnVar：0 表示成功，1 表示失败；启动或执行失败时返回空串。
func System(command string, returnVar *int) string {
	*returnVar = 0
	var stdBuf bytes.Buffer
	var err, err1, err2, err3 error

	// split command
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})
	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}
	cmd := exec.Command(parts[0], parts[1:]...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := io.MultiWriter(os.Stdout, &stdBuf)
	stderr := io.MultiWriter(os.Stderr, &stdBuf)

	err = cmd.Start()
	if err != nil {
		*returnVar = 1
		return ""
	}

	go func() {
		_, err1 = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, err2 = io.Copy(stderr, stderrIn)
	}()

	err3 = cmd.Wait()
	if err1 != nil || err2 != nil || err3 != nil {
		if err1 != nil {
			fmt.Println(err1)
		}
		if err2 != nil {
			fmt.Println(err2)
		}
		if err3 != nil {
			fmt.Println(err3)
		}
		*returnVar = 1
		return ""
	}
	if output := strings.TrimRight(stdBuf.String(), "\n"); output != "" {
		pos := strings.LastIndex(output, "\n")
		if pos == -1 {
			return output
		}
		return output[pos+1:]
	}
	return ""
}

// Passthru 实现 PHP passthru() 的语义：直接把命令输出原样转发到当前进程的 stdout/stderr。
// returnVar：0 表示成功，1 表示失败；失败时会打印错误信息。
func Passthru(command string, returnVar *int) {
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})
	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		*returnVar = 1
		fmt.Println(err)
	} else {
		*returnVar = 0
	}
}

// Echo 实现 PHP echo：将参数依次输出到标准输出，不带换行。
func Echo(args ...interface{}) {
	fmt.Print(args...)
}

// Uniqid 实现 PHP uniqid(prefix)：生成带前缀的“唯一 ID”，由 Unix 秒 + 纳秒后 20 位组成。
func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}

// Exit 实现 PHP exit(status)：以给定状态码终止当前进程。
func Exit(status int) {
	os.Exit(status)
}

// Die 是 Exit 的别名。
func Die(status int) {
	os.Exit(status)
}

// Getenv 实现 PHP getenv(varname)：读取环境变量，不存在时返回空串。
func Getenv(varname string) string {
	return os.Getenv(varname)
}

// Putenv 实现 PHP putenv(setting)：解析 "NAME=VALUE" 形式并写入环境变量。
// setting 不包含 "=" 时会 panic。
func Putenv(setting string) error {
	s := strings.Split(setting, "=")
	if len(s) != 2 {
		panic("setting: invalid")
	}
	return os.Setenv(s[0], s[1])
}

// MemoryGetUsage 实现 PHP memory_get_usage()：返回当前分配的堆内存字节数。
// realUsage 当前未使用；保留以匹配 PHP 行为。
func MemoryGetUsage(realUsage bool) uint64 {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	return stat.Alloc
}

// VersionCompare 实现 PHP version_compare()。
//
// 支持的运算符：<, lt, <=, le, >, gt, >=, ge, ==, =, eq, !=, <>, ne。
// 特殊版本字符串按以下顺序比较：
//
//	(任意未识别) < dev < alpha = a < beta = b < RC = rc < # < pl = p
//
// 示例：
//
//	VersionCompare("1.2.3-alpha", "1.2.3RC7", ">=")
//	VersionCompare("1.2.3-beta", "1.2.3pl", "lt")
//	VersionCompare("1.1_dev", "1.2any", "eq")
//
// 非法 operator 会触发 panic。
func VersionCompare(version1, version2, operator string) bool {
	var vcompare func(string, string) int
	var canonicalize func(string) string
	var special func(string, string) int

	// version compare
	vcompare = func(origV1, origV2 string) int {
		if origV1 == "" || origV2 == "" {
			if origV1 == "" && origV2 == "" {
				return 0
			}
			if origV1 == "" {
				return -1
			}
			return 1
		}

		ver1, ver2, compare := "", "", 0
		if origV1[0] == '#' {
			ver1 = origV1
		} else {
			ver1 = canonicalize(origV1)
		}
		if origV2[0] == '#' {
			ver2 = origV2
		} else {
			ver2 = canonicalize(origV2)
		}
		n1, n2 := 0, 0
		for {
			p1, p2 := "", ""
			n1 = strings.IndexByte(ver1, '.')
			if n1 == -1 {
				p1, ver1 = ver1[:], ""
			} else {
				p1, ver1 = ver1[:n1], ver1[n1+1:]
			}
			n2 = strings.IndexByte(ver2, '.')
			if n2 == -1 {
				p2, ver2 = ver2, ""
			} else {
				p2, ver2 = ver2[:n2], ver2[n2+1:]
			}

			if (p1[0] >= '0' && p1[0] <= '9') && (p2[0] >= '0' && p2[0] <= '9') { // all is digit
				l1, _ := strconv.Atoi(p1)
				l2, _ := strconv.Atoi(p2)
				if l1 > l2 {
					compare = 1
				} else if l1 == l2 {
					compare = 0
				} else {
					compare = -1
				}
			} else if !(p1[0] >= '0' && p1[0] <= '9') && !(p2[0] >= '0' && p2[0] <= '9') { // all digit
				compare = special(p1, p2)
			} else { // part is digit
				if p1[0] >= '0' && p1[0] <= '9' { // is digit
					compare = special("#N#", p2)
				} else {
					compare = special(p1, "#N#")
				}
			}

			if compare != 0 || n1 == -1 || n2 == -1 {
				break
			}
		}

		if compare == 0 {
			if ver1 != "" {
				if ver1[0] >= '0' && ver1[0] <= '9' {
					compare = 1
				} else {
					compare = vcompare(ver1, "#N#")
				}
			} else if ver2 != "" {
				if ver2[0] >= '0' && ver2[0] <= '9' {
					compare = -1
				} else {
					compare = vcompare("#N#", ver2)
				}
			}
		}

		return compare
	}

	// canonicalize
	canonicalize = func(version string) string {
		ver := []byte(version)
		l := len(ver)
		if l == 0 {
			return ""
		}
		buf := make([]byte, l*2)
		j := 0
		for i, v := range ver {
			next := uint8(0)
			if i+1 < l { // Have the next one
				next = ver[i+1]
			}
			if v == '-' || v == '_' || v == '+' { // replace '-', '_', '+' to '.'
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else if (next > 0) &&
				(!(next >= '0' && next <= '9') && (v >= '0' && v <= '9')) ||
				(!(v >= '0' && v <= '9') && (next >= '0' && next <= '9')) { // Insert '.' before and after a non-digit
				buf[j] = v
				j++
				if v != '.' && next != '.' {
					buf[j] = '.'
					j++
				}
				continue
			} else if !((v >= '0' && v <= '9') ||
				(v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')) { // Non-letters and numbers
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else {
				buf[j] = v
				j++
			}
		}

		return string(buf[:j])
	}

	// compare special version forms
	special = func(form1, form2 string) int {
		found1, found2, len1, len2 := -1, -1, len(form1), len(form2)
		// (Any string not found) < dev < alpha = a < beta = b < RC = rc < # < pl = p
		forms := map[string]int{
			"dev":   0,
			"alpha": 1,
			"a":     1,
			"beta":  2,
			"b":     2,
			"RC":    3,
			"rc":    3,
			"#":     4,
			"pl":    5,
			"p":     5,
		}

		for name, order := range forms {
			if len1 < len(name) {
				continue
			}
			if strings.Compare(form1[:len(name)], name) == 0 {
				found1 = order
				break
			}
		}
		for name, order := range forms {
			if len2 < len(name) {
				continue
			}
			if strings.Compare(form2[:len(name)], name) == 0 {
				found2 = order
				break
			}
		}

		if found1 == found2 {
			return 0
		} else if found1 > found2 {
			return 1
		} else {
			return -1
		}
	}

	compare := vcompare(version1, version2)

	switch operator {
	case "<", "lt":
		return compare == -1
	case "<=", "le":
		return compare != 1
	case ">", "gt":
		return compare == 1
	case ">=", "ge":
		return compare != -1
	case "==", "=", "eq":
		return compare == 0
	case "!=", "<>", "ne":
		return compare != 0
	default:
		panic("operator: invalid")
	}
}

// ZipOpen 打开一个 zip 文件并返回 *zip.ReadCloser。
func ZipOpen(filename string) (*zip.ReadCloser, error) {
	return zip.OpenReader(filename)
}

// Pack 按指定的字节序 order 将 data 编码为字符串。
// 写入失败时返回错误；data 必须为定长类型。
func Pack(order binary.ByteOrder, data interface{}) (string, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Unpack 是 Pack 的逆操作：从 data 字符串按字节序 order 解码到 []byte。
func Unpack(order binary.ByteOrder, data string) (interface{}, error) {
	var result []byte
	r := bytes.NewReader([]byte(data))
	err := binary.Read(r, order, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Ternary 实现三元表达式：根据 condition 返回 trueVal 或 falseVal。
// 调用方需自行做类型断言，例如 max := Ternary(a > b, a, b).(int)。
func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
