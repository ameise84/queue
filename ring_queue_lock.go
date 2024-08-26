package queue

import (
	"sync"
)

func NewRingQueueLock[T any](cap int32, mu sync.Locker) *RingQueueLock[T] {
	q := &RingQueueLock[T]{q: NewRingQueue[T](cap)}
	q.mu = mu
	return q
}

type RingQueueLock[T any] struct {
	q  *RingQueue[T]
	mu sync.Locker
}

func (r *RingQueueLock[T]) ResetCap(cap int32) error {
	r.mu.Lock()
	err := r.q.ResetCap(cap)
	r.mu.Unlock()
	return err
}

func (r *RingQueueLock[T]) PushFront(v T) error {
	r.mu.Lock()
	ok := r.q.PushFront(v)
	r.mu.Unlock()
	return ok
}

func (r *RingQueueLock[T]) PushBack(v T) error {
	r.mu.Lock()
	ok := r.q.PushBack(v)
	r.mu.Unlock()
	return ok
}

func (r *RingQueueLock[T]) PopFront() (T, error) {
	r.mu.Lock()
	v, ok := r.q.PopFront()
	r.mu.Unlock()
	return v, ok
}

func (r *RingQueueLock[T]) PopBack() (T, error) {
	r.mu.Lock()
	v, ok := r.q.PopBack()
	r.mu.Unlock()
	return v, ok
}

func (r *RingQueueLock[T]) Enqueue(v T) error {
	return r.PushBack(v)
}

func (r *RingQueueLock[T]) Dequeue() (T, error) {
	return r.PopFront()
}

func (r *RingQueueLock[T]) IsEmpty() bool {
	r.mu.Lock()
	b := r.q.IsEmpty()
	r.mu.Unlock()
	return b
}
