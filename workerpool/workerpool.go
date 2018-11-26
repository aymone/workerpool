package workerpool

import (
	"fmt"
)

// Job interface
type Job interface {
	Do()
}

type pool struct {
	name string
	size int
	jobs chan Job
}

// New starts a pool of workers
func New(name string, poolSize int) *pool {
	p := &pool{
		name: name,
		size: poolSize,
		jobs: make(chan Job),
	}

	go p.process()

	return p
}

func (p *pool) AddJob(j Job) {
	p.jobs <- j
}

func (p *pool) process() {
	tickets := make(chan bool, p.size)
	fmt.Printf("\nProcessing pool '%s' with %d workers", p.name, p.size)
	for j := range p.jobs {
		tickets <- true
		go func(j Job, t chan bool) {
			j.Do()
			<-t
		}(j, tickets)
	}
}
