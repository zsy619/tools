package queues

import (
	"sync"

	"haedu.gov.cn/tools/xgeneric/lists"
)

// Queue provides a container as the rule of FIFO(first in first out) manner
type Queue[E any] struct {
	l    *lists.List[*QueueEle[E]]
	lock sync.Mutex

	cond *sync.Cond
}

type QueueEle[E any] struct {
	v E
}

// NewQueue create a new queue
func NewQueue[E any]() *Queue[E] {
	return &Queue[E]{l: lists.NewList[*QueueEle[E]](), cond: sync.NewCond(&sync.Mutex{})}
}

// Enqueue add element to queue
func (q *Queue[E]) Enqueue(e E) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.l.PushBack(&QueueEle[E]{e})
	q.cond.Signal()
}

// Dequeue dequeue element from top
func (q *Queue[E]) Dequeue() E {
	var ret E
	e := q.dequeueEle()
	if e != nil {
		ret = e.v
	}

	return ret
}

// Dequeue dequeue element from top
func (q *Queue[E]) dequeueEle() *QueueEle[E] {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.l.RemoveFront()
}

// Clear all elements
func (q *Queue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.l.Clear()
}
