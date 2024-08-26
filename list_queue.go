package queue

import (
	"sync"
)

var listQueueElementPool sync.Pool

func init() {
	listQueueElementPool = sync.Pool{New: func() any {
		return &Element{}
	}}
}

type Element struct {
	value any
	prev  *Element
	next  *Element
}

func NewListQueue[T any]() *ListQueue[T] {
	n := listQueueElementPool.Get().(*Element)
	n.prev = nil
	n.next = nil
	return &ListQueue[T]{head: n, tail: n}
}

type ListQueue[T any] struct {
	head *Element
	tail *Element
	zero T
}

func (l *ListQueue[T]) ResetCap(uint32) (uint32, error) {
	panic("can not reset capacity")
}

func (l *ListQueue[T]) PushFront(v T) error {
	newHead := listQueueElementPool.Get().(*Element)
	newHead.value = v

	newHead.next = l.head
	l.head.prev = newHead

	l.head = newHead
	return nil
}

func (l *ListQueue[T]) PushBack(v T) error {
	newTail := listQueueElementPool.Get().(*Element)
	l.tail.value = v

	newTail.prev = l.tail
	l.tail.next = newTail

	l.tail = newTail
	return nil
}

func (l *ListQueue[T]) PopFront() (T, error) {
	if l.head == l.tail {
		return l.zero, ErrQueueIsEmpty
	}
	head := l.head
	l.head = head.next
	v := head.value

	head.value = nil
	head.next = nil
	listQueueElementPool.Put(head)
	return v.(T), nil
}

func (l *ListQueue[T]) PopBack() (T, error) {
	if l.head == l.tail {
		return l.zero, ErrQueueIsEmpty
	}

	tail := l.tail.prev
	if tail == l.head {
		l.head = l.tail
	} else {
		tail.prev.next = l.tail
	}
	l.tail.prev = tail.prev

	v := tail.value

	tail.value = nil
	tail.prev = nil
	tail.next = nil
	listQueueElementPool.Put(tail)
	return v.(T), nil
}

func (l *ListQueue[T]) Enqueue(v T) error {
	return l.PushBack(v)
}

func (l *ListQueue[T]) Dequeue() (T, error) {
	return l.PopFront()
}

func (l *ListQueue[T]) IsEmpty() bool {
	return l.head.next == nil
}
