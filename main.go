package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/aymone/workerpool/job"
	"github.com/aymone/workerpool/service"
	"github.com/aymone/workerpool/workerpool"
)

func main() {
	const poolSize = 3

	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}

	// pool
	p := workerpool.New(ctx, wg, poolSize)

	// service
	s := service.New()

	counter := 0

	go func() {
		for {
			fmt.Print("\nChecking jobs")

			if p.IsAvailable() {
				fmt.Print("\nHas available space on pool :D")

				hasJob := rand.Intn(3)
				if hasJob > 1 {
					counter++
					fmt.Printf("\nJob %d awaiting to be started", counter)

					j := job.New(counter, s)
					p.Add(j)

					fmt.Printf("\nJob %d added, total is %d", counter, p.Count())
				}

			} else {
				fmt.Print("\nHasn't available space on pool :( ")
			}

			select {
			case <-ctx.Done():
				fmt.Println("\nMain: caller has told us to stop to check new jobs")
				return
			default:
				time.Sleep(time.Second * 1)
				continue
			}
		}
	}()

	// listen for C-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("\nMain: received C-c - shutting down")

	// tell the goroutines to stop
	fmt.Println("\nMain: telling goroutines to stop")
	cancel()

	// and wait for them both to reply back
	wg.Wait()
	fmt.Println("\nMain: all goroutines have told us they've finished")
}
