package queue

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// 经过测试lock free的性能不一点快
var listLockFreeQueueElementPool sync.Pool

func init() {
	listLockFreeQueueElementPool = sync.Pool{New: func() any {
		return &listLockFreeNode{}
	}}
}

type listLockFreeNode struct {
	value any
	next  unsafe.Pointer
	_     [64 - (16 + 8)]byte
}

func NewListQueueLockFree[T any]() *ListQueueLockFree[T] {
	n := unsafe.Pointer(&listLockFreeNode{})
	return &ListQueueLockFree[T]{head: n, tail: n}
}

type ListQueueLockFree[T any] struct {
	head   unsafe.Pointer
	tail   unsafe.Pointer
	length int32
	zero   T
}

func load(p *unsafe.Pointer) (n *listLockFreeNode) {
	return (*listLockFreeNode)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *listLockFreeNode) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}

func (l *ListQueueLockFree[T]) PushBack(v T) error {
	n := &listLockFreeNode{}
	var tail *listLockFreeNode
	for {
		tail = load(&l.tail)
		//更新尾节点
		if cas(&l.tail, tail, n) {
			break
		}
		runtime.Gosched()
	}
	tail.value = v
	tail.next = unsafe.Pointer(n)
	return nil
}

func (l *ListQueueLockFree[T]) PopFront() (T, error) {
	var head *listLockFreeNode
	var next *listLockFreeNode
	for {
		head = load(&l.head)
		next = load(&head.next)
		if next == nil {
			return l.zero, ErrQueueIsEmpty
		}
		//更新头节点
		if cas(&l.head, head, next) {
			break
		}
		runtime.Gosched()
	}
	v := head.value
	head.next = nil
	listLockFreeQueueElementPool.Put(head)
	return v.(T), nil
}

func (l *ListQueueLockFree[T]) Enqueue(v T) error {
	return l.PushBack(v)
}

func (l *ListQueueLockFree[T]) Dequeue() (T, error) {
	return l.PopFront()
}

func (l *ListQueueLockFree[T]) IsEmpty() bool {
	var head *listLockFreeNode
	var next *listLockFreeNode
	head = load(&l.head)
	next = load(&head.next)
	if next == nil {
		return true
	}
	return false
}
