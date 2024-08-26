package queue

func NewChanQueue[T any](cap int) IQueue[T] {
	return &ChanQueue[T]{ch: make(chan T, cap)}
}

type ChanQueue[T any] struct {
	ch   chan T
	zero T
}

func (q *ChanQueue[T]) ResetCap(uint32) (uint32, error) {
	panic("can not reset capacity")
}

func (q *ChanQueue[T]) Enqueue(v T) error {
	for {
		select {
		case q.ch <- v:
			return nil
		default:
			return ErrQueueIsFull
		}
	}
}

func (q *ChanQueue[T]) Dequeue() (T, error) {
	for {
		select {
		case v := <-q.ch:
			return v, nil
		default:
			return q.zero, ErrQueueIsEmpty
		}
	}
}

func (q *ChanQueue[T]) IsEmpty() bool {
	return len(q.ch) == 0
}
