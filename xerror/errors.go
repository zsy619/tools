package xerror

import (
	"fmt"
	"io"
)

// PrintError 调用 input 并在返回错误时打印到标准输出。
//
// 该函数主要用于以下场景：
//   - 临时脚本或一次性任务中需要快速打印错误；
//   - 测试场景下不方便引入日志器时；
//
// 生产环境建议改用结构化日志（xlog 等）替代。
func PrintError(input func() error) {
	if err := input(); err != nil {
		fmt.Println(err.Error())
	}
}

// PrintErrorTo 与 PrintError 类似，但允许指定输出目标。
//
// 当 w 为 nil 时回退到 os.Stdout，避免空指针异常。传入 *os.File、
// *bufio.Writer、*bytes.Buffer 都可以正确工作。
//
// 示例：
//
//	xerror.PrintErrorTo(os.Stderr, func() error { return doIt() })
func PrintErrorTo(w io.Writer, input func() error) {
	if err := input(); err != nil {
		if w == nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Fprintln(w, err.Error())
	}
}
