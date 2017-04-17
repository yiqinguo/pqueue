package pqueue

import (
	"sync"
)

type node struct {
	value    interface{}
	priority int64
}

type PQueue struct {
	sync.RWMutex
	nodes []*node
	count int64
}

func newNode(value interface{}, priority int64) *node {
	return &node{
		value:    value,
		priority: priority,
	}
}

func NewPQueue() *PQueue {

	nodes := make([]*node, 1)
	nodes[0] = nil

	return &PQueue{
		nodes: nodes,
		count: 0,
	}
}

func (q *PQueue) Push(value interface{}, priority int64) {
	node := newNode(value, priority)

	q.Lock()
	defer q.Unlock()
	q.nodes = append(q.nodes, node)
	q.count += 1
	q.swim(q.size())
}

func (q *PQueue) Pop() (interface{}, int64) {
	q.Lock()
	defer q.Unlock()

	if q.size() < 1 {
		return nil, 0
	}

	var max *node = q.nodes[1]
	var s int64 = 1

	q.exch(s, q.size())
	q.nodes = q.nodes[0:q.size()]
	q.count -= s
	q.sink(s)

	return max.value, max.priority
}

func (q *PQueue) Len() int64 {
	q.RLock()
	defer q.RUnlock()
	return q.size()
}

func (q *PQueue) size() int64 {
	return q.count
}

func (q *PQueue) less(i, j int64) bool {
	return q.nodes[i].priority < q.nodes[j].priority
}

func (q *PQueue) exch(i, j int64) {
	tmpNode := q.nodes[i]

	q.nodes[i] = q.nodes[j]
	q.nodes[j] = tmpNode
}

func (q *PQueue) swim(k int64) {
	for k > 1 && q.less(k/2, k) {
		q.exch(k/2, k)
		k /= 2
	}

}

func (q *PQueue) sink(k int64) {
	for 2*k <= q.size() {
		j := k * 2

		if j < q.size() && q.less(j, j+1) {
			j++
		}

		if !q.less(k, j) {
			break
		}

		q.exch(k, j)
		k = j
	}
}
