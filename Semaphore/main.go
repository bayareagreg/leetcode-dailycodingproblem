package main

import (
	"fmt"
	"sync"
	"time"
)

//
// Design and implement a fair semaphore
// A fair semaphore must guarantee that goroutines acquire permits in the exact order they requested them.
//

type Semaphore struct {
	mu      sync.Mutex
	permits int
	// FIFO queue to ensure fairness
	waitQueue []chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	if n < 0 {
		panic("semaphore size must be >= 0")
	}
	return &Semaphore{
		permits: n,
	}
}

// return the time Acquire was entered to prove fairness
func (s *Semaphore) Acquire() time.Time {
	s.mu.Lock()
	now := time.Now()

	// Fast path: no waiters and permit available
	if s.permits > 0 && len(s.waitQueue) == 0 {
		s.permits--
		s.mu.Unlock()
		return now
	}

	// Slow path: enqueue
	ch := make(chan struct{})
	s.waitQueue = append(s.waitQueue, ch)
	s.mu.Unlock()

	// Wait until signaled
	<-ch
	return now
}

func (s *Semaphore) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// If someone is waiting, wake the oldest waiter
	if len(s.waitQueue) > 0 {
		ch := s.waitQueue[0]
		s.waitQueue = s.waitQueue[1:]
		close(ch)
		return
	}

	// Otherwise return permit to pool
	s.permits++
}

func main() {
	sem := NewSemaphore(1)

	for i := range 10 {
		go func(id int) {
			then := sem.Acquire()
			defer sem.Release()

			fmt.Println("entered:", then)
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(15 * time.Second)
}
