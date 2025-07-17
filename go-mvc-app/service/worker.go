package service

import (
	"fmt"
	"sync"
)

type Task struct {
	ID int
}

var Pool = sync.Pool{
	New: func() any {
		return &Task{}
	},
}

func StartWorkPool(jobs <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				task := Pool.Get().(*Task)
				task.ID = job
				fmt.Printf("worker %d is performing job %d", id, task.ID)
				Pool.Put(task)
			}
		}(i)
	}

}
