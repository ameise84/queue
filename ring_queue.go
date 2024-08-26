package queue

import (
	"fmt"
)

func NewRingQueue[T any](cap int32) *RingQueue[T] {
	if cap <= 0 {
		panic(fmt.Sprintf("ring-queue capacity[%d] is smaller", cap))
	}

	r := &RingQueue[T]{
		cell: make([]T, cap),
		cap:  cap,
	}
	return r
}

type RingQueue[T any] struct {
	cell    []T
	cap     int32
	size    int32
	tailIdx int32
	headIdx int32
	zero    T
}

func (r *RingQueue[T]) ResetCap(cap int32) error {
	if cap <= 0 {
		return ErrQueueCapacityZero
	}
	if r.tailIdx != r.headIdx {
		return ErrQueueIsNotEmpty
	}
	if cap > r.cap {
		r.cell = make([]T, cap)
	}
	r.cap = cap
	return nil
}

func (r *RingQueue[T]) PushFront(v T) error {
	if r.size >= r.cap {
		return ErrQueueIsFull
	}
	r.size++
	r.headIdx--
	r.fixIndex(&r.headIdx)
	r.cell[r.headIdx] = v
	return nil
}

func (r *RingQueue[T]) PushBack(v T) error {
	if r.size >= r.cap {
		return ErrQueueIsFull
	}
	r.cell[r.tailIdx] = v
	r.size++
	r.tailIdx++
	r.fixIndex(&r.tailIdx)
	return nil
}

func (r *RingQueue[T]) PopFront() (T, error) {
	if r.size == 0 {
		return r.zero, ErrQueueIsEmpty
	}
	v := r.cell[r.headIdx]
	r.size--
	r.headIdx++
	r.fixIndex(&r.headIdx)
	return v, nil
}

func (r *RingQueue[T]) PopBack() (T, error) {
	if r.size == 0 {
		return r.zero, ErrQueueIsEmpty
	}
	r.size--
	r.tailIdx--
	r.fixIndex(&r.tailIdx)
	v := r.cell[r.tailIdx]
	return v, nil
}

func (r *RingQueue[T]) Enqueue(v T) error {
	return r.PushBack(v)
}

func (r *RingQueue[T]) Dequeue() (T, error) {
	return r.PopFront()
}

func (r *RingQueue[T]) IsEmpty() bool {
	return r.size == 0
}

func (r *RingQueue[T]) fixIndex(index *int32) {
	if *index < 0 {
		*index = r.cap + *index
	} else if *index >= r.cap {
		*index = *index - r.cap
	}
}
