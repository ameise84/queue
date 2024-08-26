package queue

import (
	"sync"
)

func NewListQueueLock[T any](mu sync.Locker) *ListQueueLock[T] {
	q := &ListQueueLock[T]{q: NewListQueue[T]()}
	q.mu = mu
	return q
}

type ListQueueLock[T any] struct {
	q  *ListQueue[T]
	mu sync.Locker
}

func (l *ListQueueLock[T]) PushFront(v T) error {
	l.mu.Lock()
	err := l.q.PushFront(v)
	l.mu.Unlock()
	return err
}

func (l *ListQueueLock[T]) PushBack(v T) error {
	l.mu.Lock()
	err := l.q.PushBack(v)
	l.mu.Unlock()
	return err
}

func (l *ListQueueLock[T]) PopFront() (T, error) {
	l.mu.Lock()
	v, err := l.q.PopFront()
	l.mu.Unlock()
	return v, err
}

func (l *ListQueueLock[T]) PopBack() (T, error) {
	l.mu.Lock()
	v, err := l.q.PopBack()
	l.mu.Unlock()
	return v, err
}

func (l *ListQueueLock[T]) Enqueue(v T) error {
	return l.PushBack(v)
}

func (l *ListQueueLock[T]) Dequeue() (T, error) {
	return l.PopFront()
}

func (l *ListQueueLock[T]) IsEmpty() bool {
	l.mu.Lock()
	ok := l.q.IsEmpty()
	l.mu.Unlock()
	return ok
}
