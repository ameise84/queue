package queue

type IQueue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	IsEmpty() bool
}

type IDeque[T any] interface {
	PushFront(T) error
	PushBack(T) error
	PopFront() (T, error)
	PopBack() (T, error)
	IsEmpty() bool
}
