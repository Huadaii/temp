package queue

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the queue
type Item generic.Type

type ItemQueue struct {
	Data []Item
	Mux  sync.RWMutex
}

func (s *ItemQueue) New() *ItemQueue {
	s.Data = []Item{}
	s.Mux = sync.RWMutex{}
	return s
}

func (s *ItemQueue) Enqueue(item Item) {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	s.Data = append(s.Data, item)
}

func (s *ItemQueue) Dequeue() Item {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	item := s.Data[0]
	s.Data = s.Data[1:]

	return item
}

func (s *ItemQueue) Waiting() Item {
	s.Mux.RLock()
	defer s.Mux.RUnlock()
	return s.Data[0 : len(s.Data)-1]
}

func (s *ItemQueue) Front() Item {
	s.Mux.RLock()
	defer s.Mux.RUnlock()
	return s.Data[0]
}

func (s *ItemQueue) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *ItemQueue) Size() int {
	return len(s.Data)
}
