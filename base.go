// Package tools 是 zsy/tools 工具库的根包，提供一系列基础类型别名、
// 函数式接口与通用辅助能力，作为其它子包（xarray、xstring、xmath…）
// 的共同依赖。
//
// 该包有意保持轻量，仅暴露与具体业务无关的通用抽象，例如函数类型别名、
// 可比较接口以及一个表示 "无值" 标记的 Null 占位类型，避免在多个子包中
// 重复定义相同概念。
//
// 命名规范：
//   - 以 Null、Call、Empty 等导出符号代表语义占位或回调入口；
//   - 以 Func/BiFunc/Consumer/Supplier 等代表 Java 风格或函数式编程
//     领域常见的函数类型别名；
//   - 以 Evaluate/CMP/EQL/Comparable 等代表可复用的判定、比较接口。
//
// 该包不依赖任何第三方模块，也不参与运行时调度，仅作为静态类型集合存在。
package tools

type (
	// Null 是一个零字节结构体，用作 "无值" 或 "占位" 的语义标记。
	//
	// 典型用途：在需要表示 "与 nil 不同但仍可视为空" 的场景下使用，
	// 例如在容器 API 中表示 "没有元素" 但又不想直接返回 nil 通道
	// 或空切片。Null 与空结构体共享地址，可以放心地大量分配而几乎
	// 不会带来内存开销。
	Null struct{}

	// Call 表示一个无参数无返回值的回调函数签名。
	//
	// 用于包装需要在某处延迟执行的副作用，例如钩子、生命周期回调等。
	Call func()

	// Func 表示一个接收一个参数 T 并返回类型 R 结果的函数。
	//
	// 等价于 func(T) R，是函数式编程中 map/filter/reduce 等基础
	// 操作的最常见签名。
	Func[R, T any] func(T) R

	// BiFunc 是 Func 的二元版本：接收两个参数 T、U，并返回类型 R 的结果。
	//
	// 适用于双参数的归约（reduce）、键值对遍历等场景。
	BiFunc[R, T, U any] func(T, U) R

	// Consumer 表示接收一个参数 T 但不返回任何结果的函数。
	//
	// 常用于 forEach、订阅通知、写入等只关心副作用的回调。
	Consumer[T any] func(T)

	// BiConsumer 是 Consumer 的二元版本：接收两个参数 T、U 但不返回结果。
	BiConsumer[T, U any] func(T, U)

	// Supplier 表示不接受任何参数但返回一个 R 类型结果的工厂函数。
	//
	// 用于惰性求值、对象工厂、延迟初始化等场景。
	Supplier[R any] func() R

	// Evaluate 表示使用指定参数 E 执行测试，并返回布尔结果的判定函数。
	//
	// 可理解为 filter/Predicate 在强类型语言中的等价物。
	Evaluate[E any] Func[bool, E]

	// CMP 表示比较函数：接收两个同类型元素，返回 int。
	//
	// 返回值遵循惯例：负值表示 a<b，零表示相等，正值表示 a>b。
	// 与标准库 cmp.Compare 保持一致的语义。
	CMP[E any] BiFunc[int, E, E]

	// EQL 表示相等判定函数：接收两个同类型元素，返回 bool。
	//
	// 在业务中存在与 == 不同的相等语义（例如深度比较、近似比较）时使用。
	EQL[E any] BiFunc[bool, E, E]

	// Comparable 是一个可比较接口：实现 CompareTo 即可参与排序或有序容器。
	//
	// 实现者可以是任意类型 E；返回值含义与 CMP 一致：
	//   - 负值表示当前对象小于 v；
	//   - 零表示当前对象等于 v；
	//   - 正值表示当前对象大于 v。
	Comparable[E any] interface {
		CompareTo(v E) int
	}
)

// Empty 是 Null 的零值实例，作为 "空 / 占位 / 无值" 的全局单例使用。
//
// 推荐用法：
//
//	if got == tools.Empty { /* 表示缺失 */ }
//
// 注意：Empty 仅作语义占位，并非线程安全的同步原语，使用者不应
// 对其加锁或修改（其类型是值类型，无法修改）。
var Empty Null
