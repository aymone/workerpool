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
	p := workerpool.New(ctx, wg, "My pool", poolSize)

	// service
	s := service.New()

	counter := 0

	go func() {
		for {
			fmt.Print("\nChecking jobs")

			hasJob := rand.Intn(3)
			if hasJob > 1 {
				counter++
				fmt.Printf("\njob %d awaiting to be started", counter)

				j := job.New(counter, s)
				p.Add(j)

				fmt.Printf("\njob %d added, total is %d", counter, p.Count())
			}

			time.Sleep(time.Second * 1)

			select {
			case <-ctx.Done():
				fmt.Println("tock: caller has told us to stop")
				return
			default:
				continue
			}
		}
	}()

	// listen for C-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("main: received C-c - shutting down")

	// tell the goroutines to stop
	fmt.Println("main: telling goroutines to stop")
	cancel()

	// and wait for them both to reply back
	wg.Wait()
	fmt.Println("main: all goroutines have told us they've finished")
}
