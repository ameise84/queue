package queue

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

const (
	empty = iota
	locked
	value
)

type ringQueueLockFreeNode struct {
	v     any
	state int32
}

// NewRingQueueLockFree 由于内部运算需要,所以queue的容量必须为2^n次方
func NewRingQueueLockFree[T any](cap int32) *RingQueueLockFree[T] {
	if cap <= 0 {
		panic(fmt.Sprintf("ring-queue-lock-free capacity[%d] is smaller", cap))
	}
	newCap := ceilToPowerOfTwo(uint32(cap))
	if newCap != uint32(cap) {
		panic(fmt.Sprintf("ring-queue-lock-free capacity[%d] must power 2", cap))
	}
	return &RingQueueLockFree[T]{cell: make([]ringQueueLockFreeNode, newCap), cap: int64(newCap), mask: newCap - 1}
}

type RingQueueLockFree[T any] struct {
	cap     int64
	mask    uint32
	size    int64 //转为int64是为了判断size<0方便
	tailIdx uint32
	headIdx uint32
	cell    []ringQueueLockFreeNode
	zero    T
}

func (r *RingQueueLockFree[T]) ResetCap(uint32) (uint32, error) {
	panic("can not reset capacity")
}

func (r *RingQueueLockFree[T]) Enqueue(v T) error {
	if r.size >= r.cap {
		return ErrQueueIsFull
	}
	size := atomic.AddInt64(&r.size, 1)
	if size >= r.cap {
		atomic.AddInt64(&r.size, -1)
		return ErrQueueIsFull
	}
	wIndex := atomic.AddUint32(&r.tailIdx, 1)
	node := &r.cell[wIndex&r.mask]
	for {
		if atomic.CompareAndSwapInt32(&node.state, empty, locked) {
			break
		}
		runtime.Gosched()
	}
	node.v = v
	atomic.StoreInt32(&node.state, value)
	return nil
}

func (r *RingQueueLockFree[T]) Dequeue() (T, error) {
	if r.size <= 0 {
		return r.zero, ErrQueueIsEmpty
	}
	size := atomic.AddInt64(&r.size, -1)
	if size < 0 {
		atomic.AddInt64(&r.size, 1)
		return r.zero, ErrQueueIsEmpty
	}
	rIndex := atomic.AddUint32(&r.headIdx, 1)
	node := &r.cell[rIndex&r.mask]
	for {
		if atomic.CompareAndSwapInt32(&node.state, value, locked) {
			break
		}
		runtime.Gosched()
	}
	v := node.v
	atomic.StoreInt32(&node.state, empty)
	return v.(T), nil
}

func (r *RingQueueLockFree[T]) IsEmpty() bool {
	return atomic.LoadInt64(&r.size) == 0
}
