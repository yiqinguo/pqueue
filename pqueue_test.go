package pqueue

import (
	"testing"
)

func Test_NewPQueue(t *testing.T) {
	q := NewPQueue()
	if len(q.nodes) == 1 {
		t.Logf("node len: %d", len(q.nodes))
	} else {
		t.Fatalf("node len: %d, must be 1", len(q.nodes))
	}
}

func Test_Push(t *testing.T) {
	q := NewPQueue()
	for i := int64(0); i < 5; i++ {
		q.Push(i, i)
	}

	if q.Len() == 5 {
		t.Logf("queue len: %d", q.Len())
	} else {
		t.Fatalf("queue len: %d must be 5")
	}
}

func Test_Pop(t *testing.T) {
	q := NewPQueue()
	d, p := q.Pop()
	if d == nil && p == 0 {
		t.Logf("queue len 0")
	} else {
		t.Fatalf("queue len: %d must be 0", q.Len())
	}

	for i := int64(0); i < 5; i++ {
		q.Push(i, i)
	}
	d, p = q.Pop()
	if p == 4 && d.(int64) == 4 {
		t.Logf("head is %d", p)
	} else {
		t.Fatalf("head is %d must be 4", p)
	}
}

func BenchmarkPush(b *testing.B) {
	q := NewPQueue()
	for i := 0; i < b.N; i++ {
		q.Push(i, int64(i))
	}
}

func BenchmarkPop(b *testing.B) {
	q := NewPQueue()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
