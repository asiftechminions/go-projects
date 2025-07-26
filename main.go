package main

import (
	"fmt"
	"sync"
)

// Worker pool size (tunable based on CPU)
const workerCount = 10

// Worker that processes even numbers, squares them, and sends to result channel
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		results <- num * num
	}
}

func sumOfEvenSquares(nums []int) int {
	jobs := make(chan int, 100)    // Buffered to reduce blocking
	results := make(chan int, 100) // Buffered to avoid goroutine blocking
	var wg sync.WaitGroup

	// Start a pool of workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Feed even numbers to the job channel
	go func() {
		for _, num := range nums {
			if num%2 == 0 {
				jobs <- num
			}
		}
		close(jobs)
	}()

	// Close results channel after all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	sum := 0
	for sq := range results {
		sum += sq
	}
	return sum
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	output := sumOfEvenSquares(input)
	fmt.Println("Sum of squares of even numbers:", output)
}
