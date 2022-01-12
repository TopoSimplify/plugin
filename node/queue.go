package node

import (
	"github.com/intdxdt/deque"
	"sync"
)

// @formatter:off

type Queue struct {
	sync.RWMutex
	que *deque.Deque
}

func NewQueue() *Queue {
	return &Queue{que: deque.NewDeque()}
}

func (q *Queue) Append(o Node) *Queue {
	q.Lock()
	q.que.Append(o)
	q.Unlock()
	return q
}

func (q *Queue) AppendLeft(o Node) *Queue {
	q.Lock()
	q.que.AppendLeft(o)
	q.Unlock()
	return q
}

func (q *Queue) Pop() Node {
	q.Lock()
	n := q.que.Pop().(Node)
	q.Unlock()
	return n
}

func (q *Queue) PopLeft() Node {
	q.Lock()
	n := q.que.PopLeft().(Node)
	q.Unlock()
	return n
}

func (q *Queue) Clear() *Queue {
	q.Lock()
	q.que.Clear()
	q.Unlock()
	return q
}

func (q *Queue) Size() int {
	q.RLock()
	n := q.que.Len()
	q.RUnlock()
	return n
}

func (q *Queue) First() Node {
	q.RLock()
	n := q.que.First().(Node)
	q.RUnlock()
	return n
}

func (q *Queue) Last() Node {
	q.RLock()
	n := q.que.Last().(Node)
	q.RUnlock()
	return n
}
