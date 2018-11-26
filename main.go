package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aymone/workerpool/job"
	"github.com/aymone/workerpool/service"
	"github.com/aymone/workerpool/workerpool"
)

func main() {
	const poolSize = 3

	// pool
	p := workerpool.New("My pool", poolSize)

	// service
	s := service.New()

	counter := 0
	for {
		fmt.Print("\nChecking jobs")

		hasJob := rand.Intn(3)
		if hasJob > 0 {
			counter++
			fmt.Printf("\njob %d awaiting to be started", counter)

			j := job.New(counter, s)
			p.AddJob(j)

			fmt.Printf("\njob %d added", counter)
		}

		time.Sleep(time.Second * 1)
	}
}
