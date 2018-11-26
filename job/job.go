package job

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aymone/workerpool/service"
)

// Samplejob is a implementation of workerpool Job interface
type Samplejob struct {
	ID      int
	service service.JobService
}

// New sample job service
func New(ID int, s service.JobService) *Samplejob {
	return &Samplejob{ID, s}
}

// Do Job implementation
func (j *Samplejob) Do() {
	waitSeconds := rand.Intn(15)
	time.Sleep(time.Second * time.Duration(waitSeconds))
	entity, err := j.service.Get(j.ID)
	if err != nil {
		fmt.Printf("\njob %d had error: %s", entity.ID, err.Error())
	}

	fmt.Printf("\njob %d done in %d seconds", entity.ID, waitSeconds)
}
