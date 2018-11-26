package workerpool

// Worker interface
type Worker interface {
	Do()
}

type pool struct {
	name    string
	jobs    chan Worker
	tickets chan bool
}

// New starts a pool of workers
func New(name string, poolSize int) *pool {
	p := &pool{
		name:    name,
		jobs:    make(chan Worker),
		tickets: make(chan bool, poolSize),
	}

	go p.process()

	return p
}

func (p *pool) AddJob(j Worker) {
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
			go func(j Worker, tickets chan bool) {
				j.Do()
				<-tickets
			}(job, p.tickets)
		}
	}
}
