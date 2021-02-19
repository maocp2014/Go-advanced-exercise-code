package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// start time
var startTime = time.Now()

// perform a job in the future
func worker(ctx context.Context, seconds int) {
	select {

	// if context closes, end goroutine
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%ds) killed!\n", time.Since(startTime).Seconds(), seconds)
		return // kills goroutine

	// do the job after `seconds` seconds
	case <-time.After(time.Duration(seconds) * time.Second):
		fmt.Printf("%0.2fs - worker(%ds) completed the job.\n", time.Since(startTime).Seconds(), seconds)
	}
}

// main goroutine
func main() {

	// deadline => at time, 3 seconds in the future
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	// cancel if `main` returns before the deadline
	defer cancel()

	// start worker goroutines
	go worker(ctx, 2) // execute a job after 2 seconds
	go worker(ctx, 6) // execute a job after 6 seconds
	go worker(ctx, 8) // execute a job after 8 seconds

	// sleep for 5 seconds
	time.Sleep(5 * time.Second)

	// number of active goroutines
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}