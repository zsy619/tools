package xmath

import "sync"

// WorkerOption 是用于修改 Worker 字段的函数选项（Functional Options 模式）。
type WorkerOption func(*Worker)

// NewWorker 创建一个新的 Worker 实例，并按顺序应用所有选项。
//
// 参数：
//   - opts: 可变数量的 WorkerOption，按传入顺序应用。
//
// 返回值：构造完成的 *Worker。
func NewWorker(opts ...WorkerOption) (worker *Worker) {
	worker = &Worker{}
	for _, opt := range opts {
		opt(worker)
	}
	return
}

// WithWorkermu 构造一个用于覆盖 Worker.mu 的选项。
//
// 参数：
//   - mu: 互斥锁（按值传入，赋值给 Worker.mu）。
//
// 返回值：应用后会将 worker.mu 置为 mu 的 WorkerOption。
func WithWorkermu(mu sync.Mutex) func(*Worker) {
	return func(worker *Worker) {
		worker.mu = mu
	}
}

// WithWorkertimestamp 构造一个用于覆盖 Worker.timestamp 的选项。
//
// 参数：
//   - timestamp: 时间戳（毫秒）。
//
// 返回值：应用后会将 worker.timestamp 置为 timestamp 的 WorkerOption。
func WithWorkertimestamp(timestamp int64) func(*Worker) {
	return func(worker *Worker) {
		worker.timestamp = timestamp
	}
}

// WithWorkerworkerId 构造一个用于覆盖 Worker.workerId 的选项。
//
// 参数：
//   - workerid: 节点 ID。
//
// 返回值：应用后会将 worker.workerId 置为 workerid 的 WorkerOption。
func WithWorkerworkerId(workerid int64) func(*Worker) {
	return func(worker *Worker) {
		worker.workerId = workerid
	}
}

// WithWorkernumber 构造一个用于覆盖 Worker.number 的选项。
//
// 参数：
//   - number: 当前毫秒内的 ID 序号。
//
// 返回值：应用后会将 worker.number 置为 number 的 WorkerOption。
func WithWorkernumber(number int64) func(*Worker) {
	return func(worker *Worker) {
		worker.number = number
	}
}
