package queues

import (
	"sync"

	"github.com/zsy619/tools/xgeneric/lists"
)

// Queue 是按 FIFO（先进先出）规则组织的并发安全队列。
// 内部基于双向链表实现，并通过 mutex与 sync.Cond 协调并发访问与阻塞等待。
type Queue[E any] struct {
	l    *lists.List[*QueueEle[E]]
	lock sync.Mutex

	cond *sync.Cond
}

// QueueEle 是队列中存储的内部元素包装类型，仅用于在链表中存放实际值 v。
type QueueEle[E any] struct {
	v E
}

// NewQueue 创建一个新的空队列。
// 返回：内部链表与条件变量均已初始化的 *Queue[E]。
func NewQueue[E any]() *Queue[E] {
	return &Queue[E]{l: lists.NewList[*QueueEle[E]](), cond: sync.NewCond(&sync.Mutex{})}
}

// Enqueue 将元素 e 加入队列尾部。
// 参数：e 为要入队的元素。
// 副作用：会在锁保护下追加元素，并唤醒一个等待 Dequeue 的消费者。
func (q *Queue[E]) Enqueue(e E) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.l.PushBack(&QueueEle[E]{e})
	q.cond.Signal()
}

// Dequeue 从队列头部取出并移除一个元素。
// 返回：队首元素的值；当队列为空时返回 T 的零值（不阻塞等待）。
func (q *Queue[E]) Dequeue() E {
	var ret E
	e := q.dequeueEle()
	if e != nil {
		ret = e.v
	}

	return ret
}

// dequeueEle 是 Dequeue 的内部实现，从链表中弹出头部元素。
func (q *Queue[E]) dequeueEle() *QueueEle[E] {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.l.RemoveFront()
}

// Clear 清空队列中的所有元素。
// 副作用：会在锁保护下移除链表中的全部节点。
func (q *Queue[E]) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.l.Clear()
}
