package workerpool

// Jobber interface
type Jobber interface {
	Do()
}

type pool struct {
	name    string
	jobs    chan Jobber
	tickets chan bool
}

// New starts a pool of workers
func New(name string, poolSize int) *pool {
	p := &pool{
		name:    name,
		jobs:    make(chan Jobber),
		tickets: make(chan bool, poolSize),
	}

	go p.process()

	return p
}

func (p *pool) AddJob(j Jobber) {
	p.jobs <- j
}

func (p *pool) CountJobs() int {
	return len(p.tickets)
}

func (p *pool) process() {
	for {
		p.tickets <- true
		select {
		case job := <-p.jobs:
			go func(j Jobber, tickets chan bool) {
				j.Do()
				<-tickets
			}(job, p.tickets)
		}
	}
}
