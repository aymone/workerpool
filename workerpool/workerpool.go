package workerpool

import (
	"context"
	"fmt"
	"sync"
)

// Worker interface
type Worker interface {
	Do()
}

type pool struct {
	ctx     context.Context
	wg      *sync.WaitGroup
	size    int
	workers chan Worker
	tickets chan bool
}

// New starts a pool of workers
func New(ctx context.Context, wg *sync.WaitGroup, poolSize int) *pool {
	p := &pool{
		ctx:     ctx,
		wg:      wg,
		size:    poolSize,
		workers: make(chan Worker),
		tickets: make(chan bool, poolSize),
	}

	go p.process()

	return p
}

func (p *pool) Add(w Worker) {
	p.workers <- w
}

func (p *pool) Count() int {
	return len(p.tickets)
}

func (p *pool) IsAvailable() bool {
	return len(p.tickets) < p.size
}

func (p *pool) process() {
	for {
		p.tickets <- true
		select {
		case worker := <-p.workers:
			if p.wg != nil {
				p.wg.Add(1)
			}

			go func(w Worker, tickets chan bool, wg *sync.WaitGroup) {
				if wg != nil {
					defer wg.Done()
				}

				w.Do()
				<-tickets
			}(worker, p.tickets, p.wg)

		case <-p.ctx.Done():
			fmt.Println("\nProcess: caller has told us to stop to get jobs")
			return
		}
	}
}
