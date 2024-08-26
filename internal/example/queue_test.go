package main

import (
	"github.com/ameise84/queue"
	"log"
	"testing"
)

func TestRingQueue(t *testing.T) {
	q := queue.NewRingQueue[int](32)
	for i := 0; i < 32; i++ {
		err := q.PushFront(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopFront()
		log.Println(item, err)
	}
	log.Println("-----------------")
	for i := 0; i < 32; i++ {
		err := q.PushFront(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopBack()
		log.Println(item, err)
	}

	log.Println("-----------------")
	for i := 0; i < 32; i++ {
		err := q.PushBack(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopFront()
		log.Println(item, err)
	}

	log.Println("-----------------")
	for i := 0; i < 32; i++ {
		err := q.PushBack(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopBack()
		log.Println(item, err)
	}
}

func TestListQueue(t *testing.T) {
	q := queue.NewListQueue[int]()
	for i := 0; i < 32; i++ {
		err := q.PushFront(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopFront()
		log.Println(item, err)
	}
	log.Println("-----------------")
	for i := 0; i < 32; i++ {
		err := q.PushFront(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopBack()
		log.Println(item, err)
	}

	log.Println("-----------------")
	for i := 0; i < 32; i++ {
		err := q.PushBack(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopFront()
		log.Println(item, err)
	}

	log.Println("-----------------")
	for i := 0; i < 32; i++ {
		err := q.PushBack(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.PopBack()
		log.Println(item, err)
	}
}

func TestChanQueue(t *testing.T) {
	q := queue.NewChanQueue[int](32)
	for i := 0; i < 32; i++ {
		err := q.Enqueue(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 32; i++ {
		item, err := q.Dequeue()
		log.Println(item, err)
	}
}
