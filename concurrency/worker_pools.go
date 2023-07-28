package concurrency

import (
	"context"
	"sync"
)

type Job interface {
	ID() string
	Do() error
}

type JobResult struct {
	JobID string
	Err   error
}

type WorkerPool interface {
	AddWorkers(count int)
	RemoveWorkers(count int)

	AddJob(ctx context.Context, job Job)            // context is used to add job to queue not for cancel of job
	Subscribe(ctx context.Context) chan<- JobResult // on context cancel subscription will expire
}

// implementation type
type workerPool struct {
	workers     []*worker
	jobQueue    chan Job
	jobResultCh chan JobResult
	mutex       sync.Mutex
}

func NewWorkerPool(n int) WorkerPool {
	wp := workerPool{
		workers:     make([]*worker, 0, n),
		jobQueue:    make(chan Job),
		jobResultCh: make(chan JobResult),
	}

	for i := 0; i < n; i++ {
		w := &worker{id: i, pool: &wp}
		wp.workers = append(wp.workers, w)
		go w.start()
	}

	return &wp
}

// add workers
func (wp *workerPool) AddWorkers(count int) {
	wp.mutex.Lock()
	defer wp.mutex.Unlock()

	for i := 0; i < count; i++ {
		w := &worker{
			id:   len(wp.workers),
			pool: wp,
		}
		wp.workers = append(wp.workers, w)
		go w.start()
	}
}

// remove workers
func (wp *workerPool) RemoveWorkers(count int) {
	wp.mutex.Lock()
	defer wp.mutex.Unlock()

	if count > len(wp.workers) {
		count = len(wp.workers) // can't exceed
	}

	for i := 0; i < count; i++ {
		w := wp.workers[len(wp.workers)-1]
		w.stop()
		wp.workers = wp.workers[:len(wp.workers)-1]
	}
}

// add job to queue
func (wp *workerPool) AddJob(ctx context.Context, job Job) {
	wp.jobQueue <- job
}

// on context cancel subscription will expire
func (wp *workerPool) Subscribe(ctx context.Context) chan<- JobResult {
	resultCh := make(chan JobResult)
	go func() {
		defer close(resultCh)

		for {
			select {
			case result := <-wp.jobResultCh:
				select {
				case resultCh <- result:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return resultCh
}

// a single worker
type worker struct {
	id   int
	pool *workerPool
	quit chan struct{}
}

func (w *worker) start() {
	go func() {
		for {
			select {
			case job := <-w.pool.jobQueue:
				// Process the job
				err := job.Do()

				// Send the job result to the result channel
				result := JobResult{
					JobID: job.ID(),
					Err:   err,
				}
				w.sendResult(result)

			case <-w.quit:
				// Worker received quit signal, exit the goroutine
				return
			}
		}
	}()
}

func (w *worker) stop() {
	// Send the quit signal to stop the worker
	w.quit <- struct{}{}
}

func (w *worker) sendResult(result JobResult) {
	// Send the job result to the result channel of the worker pool
	select {
	case <-w.quit:
		// Worker is stopped, discard the result
	case w.pool.jobResultCh <- result:
		// Result sent successfully
	}
}
