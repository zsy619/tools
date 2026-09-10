package xgeneric

// Ternary 是一行 if/else 表达式。
// 示例：https://go.dev/play/p/t-D7WBL44h2
// 参数：condition 为判断条件，ifOutput 为条件为真时返回的值，elseOutput 为条件为假时返回的值。
// 返回：根据 condition 选择 ifOutput 或 elseOutput。
func Ternary[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}

	return elseOutput
}

// TernaryF 是一行 if/else 表达式，其分支结果是函数调用得到的。
// 示例：https://go.dev/play/p/AO4VW20JoqM
// 参数：condition 为判断条件，ifFunc 为条件为真时调用的函数，elseFunc 为条件为假时调用的函数。
// 返回：调用相应函数得到的 T 类型值。仅被选中的分支函数会被调用。
func TernaryF[T any](condition bool, ifFunc func() T, elseFunc func() T) T {
	if condition {
		return ifFunc()
	}

	return elseFunc()
}

type ifElse[T any] struct {
	result T
	done   bool
}

// If 用于开启一个链式 if/else if/else 流程。
// 示例：https://go.dev/play/p/WSw3ApMxhyW
// 参数：condition 为判断条件，result 为条件为真时使用的结果。
// 返回：指向 ifElse[T] 的指针，用于继续链式调用 ElseIf、Else 等方法。
func If[T any](condition bool, result T) *ifElse[T] {
	if condition {
		return &ifElse[T]{result, true}
	}

	var t T
	return &ifElse[T]{t, false}
}

// IfF 用于开启一个链式 if/else if/else 流程，其结果通过函数延迟计算。
// 示例：https://go.dev/play/p/WSw3ApMxhyW
// 参数：condition 为判断条件，resultF 为条件为真时调用的函数（用于惰性求值）。
// 返回：指向 ifElse[T] 的指针，用于继续链式调用 ElseIf、Else 等方法。
func IfF[T any](condition bool, resultF func() T) *ifElse[T] {
	if condition {
		return &ifElse[T]{resultF(), true}
	}

	var t T
	return &ifElse[T]{t, false}
}

// ElseIf 用于在 If/IfF 之后追加条件分支。
// 示例：https://go.dev/play/p/WSw3ApMxhyW
// 参数：condition 为附加判断条件，result 为该条件为真时使用的值；仅当之前所有分支均未匹配时才会评估。
// 返回：自身指针，便于继续链式调用。
func (i *ifElse[T]) ElseIf(condition bool, result T) *ifElse[T] {
	if !i.done && condition {
		i.result = result
		i.done = true
	}

	return i
}

// ElseIfF 用于在 If/IfF 之后追加条件分支，其结果通过函数延迟计算。
// 示例：https://go.dev/play/p/WSw3ApMxhyW
// 参数：condition 为附加判断条件，resultF 为该条件为真时调用的函数；仅当之前所有分支均未匹配时才会评估。
// 返回：自身指针，便于继续链式调用。
func (i *ifElse[T]) ElseIfF(condition bool, resultF func() T) *ifElse[T] {
	if !i.done && condition {
		i.result = resultF()
		i.done = true
	}

	return i
}

// Else 用于结束 If/IfF 链式流程，并提供默认分支结果。
// 示例：https://go.dev/play/p/WSw3ApMxhyW
// 参数：result 为之前所有分支均未匹配时返回的值。
// 返回：链中第一个匹配分支的结果，若无匹配则返回 result。
func (i *ifElse[T]) Else(result T) T {
	if i.done {
		return i.result
	}

	return result
}

// ElseF 用于结束 If/IfF 链式流程，默认分支结果通过函数延迟计算。
// 示例：https://go.dev/play/p/WSw3ApMxhyW
// 参数：resultF 为之前所有分支均未匹配时调用的函数。
// 返回：链中第一个匹配分支的结果，若无匹配则返回调用 resultF() 得到的值。
func (i *ifElse[T]) ElseF(resultF func() T) T {
	if i.done {
		return i.result
	}

	return resultF()
}

type switchCase[T comparable, R any] struct {
	predicate T
	result    R
	done      bool
}

// Switch 是一个纯函数式的 switch/case/default 表达式。
// 示例：https://go.dev/play/p/TGbKUMAeRUd
// 参数：predicate 为待匹配的判别值。
// 返回：指向 switchCase[T, R] 的指针，用于链式调用 Case、Default 等方法。
func Switch[T comparable, R any](predicate T) *switchCase[T, R] {
	var result R

	return &switchCase[T, R]{
		predicate,
		result,
		false,
	}
}

// Case 用于在 Switch 之后追加匹配分支。
// 示例：https://go.dev/play/p/TGbKUMAeRUd
// 参数：val 为要匹配的值，result 为匹配成功时返回的结果；仅当之前所有 Case 未命中时才会评估。
// 返回：自身指针，便于继续链式调用。
func (s *switchCase[T, R]) Case(val T, result R) *switchCase[T, R] {
	if !s.done && s.predicate == val {
		s.result = result
		s.done = true
	}

	return s
}

// CaseF 用于在 Switch 之后追加匹配分支，其结果通过函数延迟计算。
// 示例：https://go.dev/play/p/TGbKUMAeRUd
// 参数：val 为要匹配的值，cb 为匹配成功时调用的函数；仅当之前所有 Case 未命中时才会评估。
// 返回：自身指针，便于继续链式调用。
func (s *switchCase[T, R]) CaseF(val T, cb func() R) *switchCase[T, R] {
	if !s.done && s.predicate == val {
		s.result = cb()
		s.done = true
	}

	return s
}

// Default 用于结束 Switch 链式流程，并提供默认分支结果。
// 示例：https://go.dev/play/p/TGbKUMAeRUd
// 参数：result 为之前所有 Case 均未匹配时返回的值。
// 返回：链中第一个匹配 Case 的结果，若无匹配则返回 result。
func (s *switchCase[T, R]) Default(result R) R {
	if !s.done {
		s.result = result
	}

	return s.result
}

// DefaultF 用于结束 Switch 链式流程，默认分支结果通过函数延迟计算。
// 示例：https://go.dev/play/p/TGbKUMAeRUd
// 参数：cb 为之前所有 Case 均未匹配时调用的函数。
// 返回：链中第一个匹配 Case 的结果，若无匹配则返回调用 cb() 得到的值。
func (s *switchCase[T, R]) DefaultF(cb func() R) R {
	if !s.done {
		s.result = cb()
	}

	return s.result
}
