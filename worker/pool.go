package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Pool struct {
	Jobs    chan Job
	Workers int
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewPool(workers int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		Jobs:    make(chan Job, 100),
		Workers: workers,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.Workers; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
}

func (p *Pool) worker(id int) {
	defer p.wg.Done()
	for {
		select {
		case job := <-p.Jobs:
			fmt.Printf("Worker %d processing job %d with data: %s\n", id, job.ID, job.Data)
			time.Sleep(1 * time.Second)
		case <-p.ctx.Done():
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

func (p *Pool) Stop() {
	p.cancel()
	p.wg.Wait()
}
