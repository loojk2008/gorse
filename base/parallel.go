package base

import (
	"gonum.org/v1/gonum/stat"
	"sync"
)

/* Parallel Schedulers */

// Parallel schedules and runs tasks in parallel. nTask is the number of tasks. nJob is
// the number of executors. worker is the executed function which passed a range of task
// IDs (begin, end).
func Parallel(nTask int, nJob int, worker func(begin, end int)) {
	var wg sync.WaitGroup
	wg.Add(nJob)
	for j := 0; j < nJob; j++ {
		go func(jobId int) {
			begin := nTask * jobId / nJob
			end := nTask * (jobId + 1) / nJob
			worker(begin, end)
			wg.Done()
		}(j)
	}
	wg.Wait()
}

// ParallelFor runs for loop in parallel.
func ParallelFor(begin, end int, worker func(i int)) {
	var wg sync.WaitGroup
	wg.Add(end - begin)
	for j := begin; j < end; j++ {
		go func(i int) {
			worker(i)
			wg.Done()
		}(j)
	}
	wg.Wait()
}

// ParallelMean schedules and runs tasks in parallel, then returns the mean of returned values.
// nJob is the number of executors. worker is the executed function which passed a range of task
// IDs (begin, end) and returns a double value.
func ParallelMean(nTask int, nJob int, worker func(begin, end int) (sum float64)) float64 {
	var wg sync.WaitGroup
	wg.Add(nJob)
	results := make([]float64, nJob)
	weights := make([]float64, nJob)
	for j := 0; j < nJob; j++ {
		go func(jobId int) {
			begin := nTask * jobId / nJob
			end := nTask * (jobId + 1) / nJob
			size := end - begin
			results[jobId] = worker(begin, end) / float64(size)
			weights[jobId] = float64(size) / float64(nTask)
			wg.Done()
		}(j)
	}
	wg.Wait()
	return stat.Mean(results, weights)
}
