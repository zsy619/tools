package xmath

import "sync"

type WorkerOption func(*Worker)

func NewWorker(opts ...WorkerOption) (worker *Worker) {
	worker = &Worker{}
	for _, opt := range opts {
		opt(worker)
	}
	return
}

func WithWorkermu(mu sync.Mutex) func(*Worker) {
	return func(worker *Worker) {
		worker.mu = mu
	}
}

func WithWorkertimestamp(timestamp int64) func(*Worker) {
	return func(worker *Worker) {
		worker.timestamp = timestamp
	}
}

func WithWorkerworkerId(workerid int64) func(*Worker) {
	return func(worker *Worker) {
		worker.workerId = workerid
	}
}

func WithWorkernumber(number int64) func(*Worker) {
	return func(worker *Worker) {
		worker.number = number
	}
}
