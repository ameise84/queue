package queue

import (
	"context"
	"time"
)

func NewChanQueueBlock[T any](cap int, ttl time.Duration) IQueue[T] {
	return &ChanQueueBlock[T]{ch: make(chan T, cap), ttl: ttl}
}

type ChanQueueBlock[T any] struct {
	ch   chan T
	ttl  time.Duration
	zero T
}

func (q *ChanQueueBlock[T]) ResetCap(uint32) (uint32, error) {
	panic("can not reset capacity")
}

func (q *ChanQueueBlock[T]) Enqueue(v T) error {
	if q.ttl > 0 {
		ctx, c := context.WithTimeout(context.Background(), q.ttl)
		defer c()
		select {
		case q.ch <- v:
			return nil
		case <-ctx.Done():
			return ErrQueueIsFull
		}
	} else {
		q.ch <- v
		return nil
	}
}

func (q *ChanQueueBlock[T]) Dequeue() (T, error) {
	if q.ttl > 0 {
		ctx, c := context.WithTimeout(context.Background(), q.ttl)
		defer c()
		select {
		case v := <-q.ch:
			return v, nil
		case <-ctx.Done():
			return q.zero, ErrQueueIsEmpty
		}
	} else {
		v := <-q.ch
		return v, nil
	}
}

func (q *ChanQueueBlock[T]) IsEmpty() bool {
	return len(q.ch) == 0
}
