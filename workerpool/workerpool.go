package workerpool

// Worker interface
type Worker interface {
	Do()
}

type pool struct {
	name    string
	workers chan Worker
	tickets chan bool
}

// New starts a pool of workers
func New(name string, poolSize int) *pool {
	p := &pool{
		name:    name,
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

func (p *pool) process() {
	for {
		p.tickets <- true
		select {
		case worker := <-p.workers:
			go func(w Worker, tickets chan bool) {
				w.Do()
				<-tickets
			}(worker, p.tickets)
		}
	}
}
