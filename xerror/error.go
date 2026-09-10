// Package xerror 提供一组轻量的错误处理工具，目标是让日常错误包装、
// 聚合与短路逻辑更简洁、更易复用。
//
// 设计与使用建议：
//   - WrapError 用于在原有错误之上附加上下文，适合日志与链路追踪场景；
//   - ReturnFirstError/OneByOneUntilError 用于聚合多个并行或顺序操作
//     的错误，避免业务层到处复制 for 循环；
//   - errors.go 中提供的 PrintError 主要用于测试或脚本场景。
//
// 该包有意不引入任何第三方依赖，仅依赖标准库 errors / fmt。
package xerror

import (
	"errors"
	"fmt"
)

// WrapError 将 err 用一段可格式化消息包装起来。
//
// 当 err 为 nil 时直接返回 nil，便于调用方无脑包装而无需提前判断。
// msg 与 args 的处理方式与 fmt.Sprintf 完全一致，可以使用 %w、%v、
// %s 等占位符；但需要 %w 保留错误链时，建议改用 fmt.Errorf("%w", err)
// 自行拼接。
//
// 示例：
//
//	if err := WrapError(repo.Save(user), "保存用户 %d 失败", user.ID); err != nil {
//	    return err
//	}
//
// 返回的字符串格式为："<msg>, err = <err.Error()>"。
func WrapError(err error, msg string, args ...any) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s, err = %s", fmt.Sprintf(msg, args...), err)
}

// ReturnFirstError 返回 errs 中第一个非 nil 错误；若全部为 nil 则返回 nil。
//
// 适用于以下并发或并行场景：
//
//	if err := xerror.ReturnFirstError(errA, errB, errC); err != nil { ... }
//
// 注意：调用方应当负责并发地填充 errs，本函数只做短路求值，自身并不
// 启动任何 goroutine。
func ReturnFirstError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// OneByOneUntilError 按顺序逐个执行 fns 中的回调，任一返回非 nil 错误即中止。
//
// 该函数适合用于必须按顺序进行的资源初始化、清理钩子、连接释放等场景，
// 与 defer 配合可在 panic 时通过传入恢复函数来收集错误。
//
// 示例：
//
//	if err := xerror.OneByOneUntilError(
//	    openConn,
//	    migrate,
//	    seed,
//	); err != nil {
//	    return err
//	}
func OneByOneUntilError(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

// Is 标准库别名，方便外部在使用本包时不必再额外导入 "errors"。
//
// 用于判断目标错误链中是否包含指定 sentinel 错误，行为完全等同于
// errors.Is(target, sentinel)。
func Is(err, target error) bool { return errors.Is(err, target) }

// As 标准库别名，方便外部在使用本包时不必再额外导入 "errors"。
//
// 用于将错误链上第一个匹配 target 类型的错误赋值给 target，行为完全
// 等同于 errors.As(err, target)。
func As(err error, target any) bool { return errors.As(err, target) }
