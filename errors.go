package queue

import (
	"errors"
)

var (
	ErrQueueIsFull       = errors.New("queue is full")
	ErrQueueIsEmpty      = errors.New("queue is empty")
	ErrQueueIsNotEmpty   = errors.New("queue is not empty")
	ErrQueueCapacityZero = errors.New("can not reset capacity to smaller than zero")
)
